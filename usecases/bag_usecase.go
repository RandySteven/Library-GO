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
	"log"
	"sync"
)

type bagUsecase struct {
	bagRepo  repositories_interfaces.BagRepository
	bookRepo repositories_interfaces.BookRepository
	userRepo repositories_interfaces.UserRepository
}

func (b *bagUsecase) setTx(ctx context.Context) {
	tx := b.bagRepo.GetTx(ctx)
	b.userRepo.SetTx(tx)
	b.bookRepo.SetTx(tx)
}

func (b *bagUsecase) refreshTx(ctx context.Context) {
	b.bagRepo.SetTx(nil)
	b.setTx(ctx)
}

func (b *bagUsecase) AddBookToBag(ctx context.Context, request *requests.BagRequest) (result *responses.AddBagResponse, customErr *apperror.CustomError) {
	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
	)
	userId := ctx.Value(enums.UserID).(uint64)
	bag := &models.Bag{
		BookID: request.BookID,
		UserID: userId,
	}
	if err := b.bagRepo.BeginTx(ctx); err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to begin tx`, err)
	}
	defer func() {
		defer b.refreshTx(ctx)
		if r := recover(); r != nil {
			_ = b.bagRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = b.bagRepo.RollbackTx(ctx)
			return
		} else {
			if err := b.bagRepo.CommitTx(ctx); err != nil {
				log.Println("failed to commit transaction")
				return
			}
			return
		}
	}()
	b.setTx(ctx)
	wg.Add(2)
	go func() {
		defer wg.Done()
		//1. check if book status avail
		book, err := b.bookRepo.FindByID(ctx, bag.BookID)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to find book`, err)
			return
		}
		if book.Status != enums.Available {
			customErrCh <- apperror.NewCustomError(apperror.ErrBadRequest, `book already not available`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		exist, err := b.bagRepo.CheckBagExists(ctx, bag)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to check book bag`, err)
			return
		}
		if exist {
			customErrCh <- apperror.NewCustomError(apperror.ErrBadRequest, `book already exist`, err)
			return
		}
	}()

	go func() {
		wg.Wait()
		close(customErrCh)
	}()

	if customErr = <-customErrCh; customErr != nil {
		return nil, customErr
	}

	bag, err := b.bagRepo.Save(ctx, bag)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save book`, err)
	}
	result = &responses.AddBagResponse{
		BookID: bag.BookID,
	}
	return result, nil
}

func (b *bagUsecase) GetUserBag(ctx context.Context) (result *responses.GetAllBagsResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	bagBooks, err := b.bagRepo.FindBagByUser(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find bag`, err)
	}
	bookIds := []uint64{}
	for _, bagBook := range bagBooks {
		bookIds = append(bookIds, bagBook.BookID)
	}
	books, err := b.bookRepo.FindSelectedBooksId(ctx, bookIds)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find bags`, err)
	}
	result = &responses.GetAllBagsResponse{}
	for _, book := range books {
		result.Books = append(result.Books, &responses.BookBagResponse{
			ID:    book.ID,
			Title: book.Title,
			Image: book.Image,
		})
	}
	return result, nil
}

func (b *bagUsecase) DeleteBookFromBag(ctx context.Context, request *requests.DeleteBookBagRequest) (customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	bagId := request.BookIDs
	err := b.bagRepo.DeleteByUserAndSelectedBooks(ctx, userId, bagId)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, `failed to delete user book in bag`, err)
	}
	return nil
}

func (b *bagUsecase) EmptyBag(ctx context.Context) (customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	err := b.bagRepo.DeleteUserBag(ctx, userId)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, `failed to delete user book bag`, err)
	}
	return nil
}

var _ usecases_interfaces.BagUsecase = &bagUsecase{}

func newBagUsecase(
	bagRepo repositories_interfaces.BagRepository,
	bookRepo repositories_interfaces.BookRepository,
	userRepo repositories_interfaces.UserRepository) *bagUsecase {
	return &bagUsecase{
		bagRepo:  bagRepo,
		bookRepo: bookRepo,
		userRepo: userRepo,
	}
}
