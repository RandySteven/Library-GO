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

func (b *bagUsecase) refreshTx(ctx context.Context) {
	tx := b.bagRepo.GetTx(ctx)
	b.userRepo.SetTx(tx)
	b.bookRepo.SetTx(tx)
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
		defer b.bagRepo.SetTx(nil)
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
	b.refreshTx(ctx)
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

	//3. insert book into user bag
	bag, err := b.bagRepo.Save(ctx, bag)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save book`, err)
	}
	result.BookID = bag.BookID
	return result, nil
}

func (b *bagUsecase) GetUserBag(ctx context.Context) (result []*responses.GetAllBagsResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (b *bagUsecase) DeleteBookFromBag(ctx context.Context, request *requests.BagRequest) (customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (b *bagUsecase) EmptyBag(ctx context.Context) (customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
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
