package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
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
}

func (b *borrowUsecase) refreshTx(ctx context.Context) {
	tx := b.borrowRepo.GetTx(ctx)
	b.bagRepo.SetTx(tx)
	b.bookRepo.SetTx(tx)
	b.borrowDetailRepo.SetTx(tx)
	b.userRepo.SetTx(tx)
	b.authorRepo.SetTx(tx)
	b.genreRepo.SetTx(tx)
}

func (b *borrowUsecase) BorrowTransaction(ctx context.Context) (result *responses.BorrowResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	bookIds := []uint64{}
	var (
		err error
		//wg          sync.WaitGroup
		//customErrCh = make(chan *apperror.CustomError)
	)

	if err = b.borrowRepo.BeginTx(ctx); err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to init trx`, err)
	}
	defer func() {
		if r := recover(); r != nil {
			_ = b.borrowRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = b.borrowRepo.RollbackTx(ctx)
		} else if err = b.borrowRepo.CommitTx(ctx); err != nil {
			log.Println("failed to commit transaction:", err)
		}
	}()
	b.refreshTx(ctx)

	//3. Get all books from user bag
	bagBooks, err := b.bagRepo.FindBagByUser(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get bag`, err)
	}
	for _, book := range bagBooks {
		bookIds = append(bookIds, book.BookID)
	}

	//4. Validate for each book if book is in available status
	for _, bookId := range bookIds {
		isExist, err := b.bookRepo.FindBookStatus(ctx, bookId, enums.Available)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get bag status`, err)
		}
		if isExist == false {
			return nil, apperror.NewCustomError(apperror.ErrBadRequest, `failed to get status`, err)
		}
	}

	//5. Create borrow header
	borrow := &models.Borrow{
		UserID:          userId,
		BorrowReference: utils.GenerateBorrowReference(24),
	}
	borrow, err = b.borrowRepo.Save(ctx, borrow)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save borrow`, err)
	}
	//6. Create borrow detail and update book
	for _, bookId := range bookIds {
		borrowDetail := &models.BorrowDetail{
			BorrowID:     borrow.ID,
			BookID:       bookId,
			ReturnedDate: time.Now().Add(7 * 24 * time.Hour),
		}
		borrowDetail, err = b.borrowDetailRepo.Save(ctx, borrowDetail)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save detail`, err)
		}
		err = b.bookRepo.UpdateBookStatus(ctx, bookId, enums.ReadyToTake)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save status`, err)
		}
	}

	//8. delete book from bag based on user
	err = b.bagRepo.DeleteUserBag(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to delete user bag`, err)
	}

	result = &responses.BorrowResponse{
		ID:           borrow.ID,
		BorrowID:     borrow.BorrowReference,
		UserID:       userId,
		TotalItems:   uint64(len(bookIds)),
		Status:       "SUCCESS",
		BorrowedDate: time.Now().Local(),
		ReturnedDate: time.Now().Local().Add(7 * 24 * time.Hour),
	}

	return result, nil
}

func (b *borrowUsecase) GetAllBorrows(ctx context.Context) (result []*responses.BorrowListResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
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

	if ctx.Value(enums.UserID).(uint64) != borrow.UserID {
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
			ID: id,
			User: struct {
				ID   uint64 `json:"id"`
				Name string `json:"name"`
			}{ID: user.ID, Name: user.Name},
			BorrowedBooks: bookDetailRes,
		}
		return result, nil
	}
}

func (b *borrowUsecase) ReturnBorrowBook(ctx context.Context, request *requests.ReturnRequest) (result *responses.ReturnBooksResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

var _ usecases_interfaces.BorrowUsecase = &borrowUsecase{}

func newBorrowUsecase(
	bagRepo repositories_interfaces.BagRepository,
	bookRepo repositories_interfaces.BookRepository,
	borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	userRepo repositories_interfaces.UserRepository,
	authorRepo repositories_interfaces.AuthorRepository,
	genreRepo repositories_interfaces.GenreRepository) *borrowUsecase {
	return &borrowUsecase{
		bagRepo:          bagRepo,
		bookRepo:         bookRepo,
		borrowRepo:       borrowRepo,
		borrowDetailRepo: borrowDetailRepo,
		userRepo:         userRepo,
		authorRepo:       authorRepo,
		genreRepo:        genreRepo,
	}
}
