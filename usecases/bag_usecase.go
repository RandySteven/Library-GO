package usecases

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"sync"
)

type bagUsecase struct {
	bagRepo     repositories_interfaces.BagRepository
	bookRepo    repositories_interfaces.BookRepository
	userRepo    repositories_interfaces.UserRepository
	cache       caches_interfaces.BagCache
	transaction repositories_interfaces.Transaction
}

func (b *bagUsecase) AddBookToBag(ctx context.Context, request *requests.BagRequest) (result *responses.AddBagResponse, customErr *apperror.CustomError) {
	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
		err         error
	)
	userId := ctx.Value(enums.UserID).(uint64)
	bag := &models.Bag{
		BookID: request.BookID,
		UserID: userId,
	}

	customErr = b.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {
		wg.Add(2)
		go func() {
			defer wg.Done()
			//1. check if book status avail

			book, err := b.bookRepo.FindByID(ctx, bag.BookID)
			if err != nil {
				customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to find book`, err)
				return
			}
			//bookBagCache, err := b.cache.GetBookBagCache(ctx, utils.HashID(request.BookID))
			//if err != nil {
			//	if !errors.Is(err, redis.Nil) {
			//		customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `redis error`, err)
			//	}
			//	return
			//}
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
			return customErr
		}

		bag, err = b.bagRepo.Save(ctx, bag)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to save book`, err)
		}
		return nil
	})
	if customErr != nil {
		return nil, customErr
	}

	result = &responses.AddBagResponse{
		BookID: bag.BookID,
	}
	_ = b.cache.Del(ctx, fmt.Sprintf(enums.UserBagKey, utils.HashID(userId)))
	_ = b.cache.SetBookBagCache(ctx, &models.BookBagCache{
		BookID: utils.HashID(bag.BookID),
		UserID: utils.HashID(userId),
		Status: enums.AtBag.ToString(),
	})
	return result, nil
}

func (b *bagUsecase) GetUserBag(ctx context.Context) (result *responses.GetAllBagsResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	bookBagList, _ := b.cache.GetUserBagCache(ctx, userId)
	if bookBagList != nil {
		result.Books = bookBagList
		return result, nil
	}
	bagBooks, err := b.bagRepo.FindBagByUser(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find bag`, err)
	}
	result = &responses.GetAllBagsResponse{}

	for _, bagBook := range bagBooks {
		//bookIds = append(bookIds, bagBook.BookID)
		result.Books = append(result.Books, &responses.BookBagResponse{
			ID:    bagBook.Book.ID,
			Title: bagBook.Book.Title,
			Image: bagBook.Book.Image,
		})
	}

	b.cache.SetUserBagCache(ctx, userId, result.Books)
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
	userRepo repositories_interfaces.UserRepository,
	cache caches_interfaces.BagCache,
	transaction repositories_interfaces.Transaction) *bagUsecase {
	return &bagUsecase{
		bagRepo:     bagRepo,
		bookRepo:    bookRepo,
		userRepo:    userRepo,
		cache:       cache,
		transaction: transaction,
	}
}
