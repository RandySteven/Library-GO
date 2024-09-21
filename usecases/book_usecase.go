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

type bookUsecase struct {
	userRepo       repositories_interfaces.UserRepository
	bookRepo       repositories_interfaces.BookRepository
	genreRepo      repositories_interfaces.GenreRepository
	authorRepo     repositories_interfaces.AuthorRepository
	authorBookRepo repositories_interfaces.AuthorBookRepository
	bookGenreRepo  repositories_interfaces.BookGenreRepository
}

func (b *bookUsecase) refreshTx(ctx context.Context) {
	tx := b.bookRepo.GetTx(ctx)
	b.userRepo.SetTx(tx)
	b.genreRepo.SetTx(tx)
	b.authorRepo.SetTx(tx)
}

func (b *bookUsecase) AddNewBook(ctx context.Context, request *requests.CreateBookRequest) (result *responses.CreateBookResponse, customErr *apperror.CustomError) {
	var (
		wg                      sync.WaitGroup
		authorIDs, genreIDs     []uint64
		authorIDChs, genreIDChs = make(chan []uint64), make(chan []uint64)
		bookCh                  = make(chan *models.Book)
		customErrCh             = make(chan *apperror.CustomError)
	)
	//1. begin tx
	err := b.bookRepo.BeginTx(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to init trx`, err)
	}
	b.refreshTx(ctx)
	defer func() {
		defer b.refreshTx(ctx)
		b.bookRepo.SetTx(nil)
		if r := recover(); r != nil {
			_ = b.bookRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = b.bookRepo.RollbackTx(ctx)
			return
		} else {
			if err = b.bookRepo.CommitTx(ctx); err != nil {
				log.Println("failed to commit transaction")
				return
			}
		}
	}()

	//2. search author and genre async and insert book
	wg.Add(3)
	go func() {
		defer wg.Done()
		for _, authorId := range request.Authors {
			author, err := b.authorRepo.FindByID(ctx, authorId)
			if err != nil {
				customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to find author`, err)
				return
			}
			authorIDs = append(authorIDs, author.ID)
		}
		authorIDChs <- authorIDs
	}()

	go func() {
		defer wg.Done()
		for _, genreId := range request.Genres {
			genre, err := b.genreRepo.FindByID(ctx, genreId)
			if err != nil {
				customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to find author`, err)
				return
			}
			genreIDs = append(genreIDs, genre.ID)
		}
		genreIDChs <- genreIDs
	}()

	go func() {
		defer wg.Done()
		book, err := b.bookRepo.Save(ctx, &models.Book{
			Title:       request.Title,
			Description: request.Description,
			Status:      enums.Available,
		})
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to create book`, err)
			return
		}
		bookCh <- book
	}()

	go func() {
		wg.Wait()
		close(customErrCh)
	}()

	select {
	case customErr = <-customErrCh:
		return nil, customErr
	default:
		wg.Add(2)
		book := <-bookCh
		go func() {
			defer wg.Done()
			authorIDs = <-authorIDChs
			for _, authorID := range authorIDs {
				_, err = b.authorBookRepo.Save(ctx, &models.AuthorBook{
					AuthorID: authorID,
					BookID:   book.ID,
				})
				if err != nil {
					customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to create author book`, err)
					return
				}
			}
		}()

		go func() {
			defer wg.Done()
			genreIDs = <-genreIDChs
			for _, genreID := range genreIDs {
				_, err = b.bookGenreRepo.Save(ctx, &models.BookGenre{
					GenreID: genreID,
					BookID:  book.ID,
				})
				if err != nil {
					customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to create book genre`, err)
					return
				}
			}
		}()

		wg.Wait()
		close(customErrCh)

		for customErr = range customErrCh {
			if customErr != nil {
				return nil, customErr
			}
		}

		result = &responses.CreateBookResponse{
			ID: book.ID,
		}
		return result, nil
	}
}

func (b *bookUsecase) GetAllBooks(ctx context.Context) (result []*responses.ListBooksResponse, customErr *apperror.CustomError) {
	return
}

func (b *bookUsecase) GetBookByID(ctx context.Context, id uint64) (result *responses.BookDetailResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.BookUsecase = &bookUsecase{}

func newBookUsecase(
	userRepo repositories_interfaces.UserRepository,
	bookRepo repositories_interfaces.BookRepository) *bookUsecase {
	return &bookUsecase{
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}
