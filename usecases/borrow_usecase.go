package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/messages"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

type borrowUsecase struct {
	bagRepo          repositories_interfaces.BagRepository
	bookRepo         repositories_interfaces.BookRepository
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository
	userRepo         repositories_interfaces.UserRepository
	authorRepo       repositories_interfaces.AuthorRepository
	genreRepo        repositories_interfaces.GenreRepository
	borrowCache      caches_interfaces.BorrowCache
	bookCache        caches_interfaces.BookCache
	pubsub           rabbitmqs_client.PubSub
	transaction      repositories_interfaces.Transaction
}

func (b *borrowUsecase) BorrowTransaction(ctx context.Context) (result *responses.BorrowResponse, customErr *apperror.CustomError) {
	var (
		err         error
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
		user        *models.User
		borrow      *models.Borrow
		bagBooks    []*models.Bag
	)
	userId := ctx.Value(enums.UserID).(uint64)

	customErr = b.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {

		user, err = b.userRepo.FindByID(ctx, userId)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrNotFound, `failed to search user`, err)
		}

		bags, _ := b.bagRepo.FindBagByUser(ctx, user.ID)
		if bags == nil {
			return apperror.NewCustomError(apperror.ErrBadRequest, `the bag already empty`, err)
		}

		bagBooks, err = b.bagRepo.FindBagByUser(ctx, userId)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to get bag`, err)
		}
		for _, bag := range bagBooks {
			if bag.Book.Status != enums.Available {
				return apperror.NewCustomError(apperror.ErrBadRequest, `book is not available`, fmt.Errorf(`book is not available`))
			}
		}

		borrow = &models.Borrow{
			UserID:          userId,
			BorrowReference: utils.GenerateBorrowReference(24),
		}
		borrow, err = b.borrowRepo.Save(ctx, borrow)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to save borrow`, err)
		}

		for _, bagBook := range bagBooks {
			wg.Add(1)
			go func(ctx context.Context, bagBook *models.Bag) {
				defer wg.Done()
				borrowDetail := &models.BorrowDetail{
					BorrowID:     borrow.ID,
					BookID:       bagBook.Book.ID,
					ReturnedDate: time.Now().Add(7 * 24 * time.Hour),
				}
				borrowDetail, err = b.borrowDetailRepo.Save(ctx, borrowDetail)
				if err != nil {
					customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to save detail`, err)
					return
				}
				err = b.bookRepo.UpdateBookStatus(ctx, bagBook.Book.ID, enums.ReadyToTake)
				if err != nil {
					customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to save status`, err)
					return
				}
			}(ctx, bagBook)
		}

		wg.Wait()
		close(customErrCh)

		for customErr = range customErrCh {
			return customErr
		}

		if err := b.bagRepo.DeleteUserBag(ctx, userId); err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to delete user bag`, err)
		}
		return nil
	})

	if err = b.pubsub.Send(ctx, "borrow_transactions", "borrow-transaction-success", &messages.EmailMessage{
		ID: uuid.NewString(),
		To: user.Email,
		Content: &messages.BorrowMessage{
			ID:              utils.HashID(borrow.ID),
			BorrowReference: borrow.BorrowReference,
		},
	}); err != nil {
		log.Fatal("error send to topic", err)
	}

	_ = b.bookCache.Del(ctx, enums.BooksKey)

	result = &responses.BorrowResponse{
		ID:           borrow.ID,
		BorrowID:     borrow.BorrowReference,
		UserID:       userId,
		TotalItems:   uint64(len(bagBooks)),
		Status:       "SUCCESS",
		BorrowedDate: time.Now().Local(),
		ReturnedDate: time.Now().Local().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	return result, nil
}

func (b *borrowUsecase) GetAllBorrows(ctx context.Context) (result []*responses.BorrowListResponse, customErr *apperror.CustomError) {
	result, err := b.borrowCache.GetMultiData(ctx)
	if errors.Is(err, redis.Nil) {
		log.Println("redis kosong")
		userId := ctx.Value(enums.UserID).(uint64)
		borrows, err := b.borrowRepo.FindByUserId(ctx, userId)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrows`, err)
		}
		for _, borrow := range borrows {
			result = append(result, &responses.BorrowListResponse{
				ID:              borrow.ID,
				BorrowReference: borrow.BorrowReference,
				BorrowedDate:    borrow.CreatedAt,
			})
		}
		err = b.borrowCache.SetMultiData(ctx, result)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert redis data`, err)
		}
	}

	return result, nil
}

func (b *borrowUsecase) GetBorrowDetail(ctx context.Context, id uint64) (result *responses.BorrowDetailResponse, customErr *apperror.CustomError) {
	var (
		wg              sync.WaitGroup
		customErrCh     = make(chan *apperror.CustomError)
		userCh          = make(chan *models.User)
		bookDetailResCh = make(chan []*responses.BorrowedBook)
	)
	borrow, err := b.borrowRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow`, err)
	}

	userId := ctx.Value(enums.UserID).(uint64)
	if userId != borrow.UserID {
		return nil, apperror.NewCustomError(apperror.ErrForbidden, `you can't access this detail`, nil)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		borrowDetails, err := b.borrowDetailRepo.FindByBorrowID(ctx, id)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow detail`, err)
			return
		}
		bookDetailRes := []*responses.BorrowedBook{}
		for _, borrowDetail := range borrowDetails {
			book, err := b.bookRepo.FindByID(ctx, borrowDetail.BookID)
			if err != nil {
				customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get book detail`, err)
				return
			}
			bookDetailRes = append(bookDetailRes, &responses.BorrowedBook{
				ID:           book.ID,
				Title:        book.Title,
				Image:        book.Image,
				BorrowedDate: borrowDetail.BorrowedDate,
				ReturnedDate: borrowDetail.ReturnedDate,
				HasReturned:  borrowDetail.VerifiedReturnDate != nil,
			})
		}
		bookDetailResCh <- bookDetailRes
		return
	}()

	go func() {
		defer wg.Done()
		user, err := b.userRepo.FindByID(ctx, borrow.UserID)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user detail`, err)
			return
		}
		userCh <- user
		return
	}()

	go func() {
		wg.Wait()
		close(userCh)
		close(customErrCh)
		close(bookDetailResCh)
	}()

	select {
	case customErr = <-customErrCh:
		return nil, customErr
	default:
		user := <-userCh
		bookDetailRes := <-bookDetailResCh
		result = &responses.BorrowDetailResponse{
			ID:              id,
			BorrowReference: borrow.BorrowReference,
			TotalItems:      len(bookDetailRes),
			User: struct {
				ID   uint64 `json:"id"`
				Name string `json:"name"`
			}{ID: user.ID, Name: user.Name},
			BorrowedBooks: bookDetailRes,
		}
		return result, nil
	}
}

func (b *borrowUsecase) BorrowConfirmation(ctx context.Context, request *requests.ConfirmBorrowRequest) (customErr *apperror.CustomError) {
	customErr = b.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {

		borrow, err := b.borrowRepo.FindByReferenceID(ctx, request.BorrowID)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow`, err)
		}

		borrowDetails, err := b.borrowDetailRepo.FindByBorrowID(ctx, borrow.ID)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow detail`, err)
		}

		for _, borrowDetail := range borrowDetails {
			err = b.bookRepo.UpdateBookStatus(ctx, borrowDetail.BookID, enums.Borrowed)
			if err != nil {
				return apperror.NewCustomError(apperror.ErrInternalServer, `failed to update borrow detail`, err)
			}
		}
		return nil
	})
	if customErr != nil {
		return customErr
	}

	_ = b.bookCache.Del(ctx, enums.BooksKey)

	return nil
}

var _ usecases_interfaces.BorrowUsecase = &borrowUsecase{}

func newBorrowUsecase(
	bagRepo repositories_interfaces.BagRepository,
	bookRepo repositories_interfaces.BookRepository,
	borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	userRepo repositories_interfaces.UserRepository,
	authorRepo repositories_interfaces.AuthorRepository,
	genreRepo repositories_interfaces.GenreRepository,
	borrowCache caches_interfaces.BorrowCache,
	bookCache caches_interfaces.BookCache,
	pubsub rabbitmqs_client.PubSub,
	transaction repositories_interfaces.Transaction) *borrowUsecase {
	return &borrowUsecase{
		bagRepo:          bagRepo,
		bookRepo:         bookRepo,
		borrowRepo:       borrowRepo,
		borrowDetailRepo: borrowDetailRepo,
		userRepo:         userRepo,
		authorRepo:       authorRepo,
		genreRepo:        genreRepo,
		borrowCache:      borrowCache,
		bookCache:        bookCache,
		pubsub:           pubsub,
		transaction:      transaction,
	}
}
