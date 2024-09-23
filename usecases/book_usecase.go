package usecases

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"sync"
)

type bookUsecase struct {
	awsClient      *aws_client.AWSClient
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
	b.authorBookRepo.SetTx(tx)
	b.bookGenreRepo.SetTx(tx)
}

func (b *bookUsecase) AddNewBook(ctx context.Context, request *requests.CreateBookRequest) (result *responses.CreateBookResponse, customErr *apperror.CustomError) {
	var (
		wg       sync.WaitGroup
		wg2      sync.WaitGroup
		errCh    = make(chan *apperror.CustomError, 1) // Buffered channels to prevent blocking
		errCh2   = make(chan *apperror.CustomError, 1)
		bookCh   = make(chan *models.Book, 1) // Buffered channel
		authorCh = make(chan []uint64, 1)     // Buffered channel
		genreCh  = make(chan []uint64, 1)     // Buffered channel
	)

	// 1. Begin transaction
	err := b.bookRepo.BeginTx(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to init transaction", err)
	}
	defer func() {
		if r := recover(); r != nil {
			_ = b.bookRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = b.bookRepo.RollbackTx(ctx)
		} else if err = b.bookRepo.CommitTx(ctx); err != nil {
			log.Println("failed to commit transaction:", err)
		}
	}()
	b.refreshTx(ctx)

	// 2. Search authors, genres, and insert book concurrently
	wg.Add(3)
	go b.findAuthors(ctx, request.Authors, authorCh, errCh, &wg)
	go b.findGenres(ctx, request.Genres, genreCh, errCh, &wg)
	go b.createBook(ctx, request, bookCh, errCh, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
	close(errCh)
	close(authorCh)
	close(genreCh)
	close(bookCh)

	// 3. Check for errors during the process
	if customErr = <-errCh; customErr != nil {
		return nil, customErr
	}

	// 4. Insert relationships (Authors and Genres)
	book := <-bookCh
	authorIDs := <-authorCh
	genreIDs := <-genreCh

	wg2.Add(2)
	go b.createAuthorBookRelations(ctx, authorIDs, book.ID, errCh2, &wg2)
	go b.createBookGenreRelations(ctx, genreIDs, book.ID, errCh2, &wg2)

	// Wait for second set of goroutines
	wg2.Wait()
	close(errCh2)

	// Check for errors from the second phase
	if customErr = <-errCh2; customErr != nil {
		return nil, customErr
	}

	// 5. Return response
	result = &responses.CreateBookResponse{
		ID: utils.HashID(book.ID),
	}
	return result, nil
}

// findAuthors fetches authors by their IDs concurrently
func (b *bookUsecase) findAuthors(ctx context.Context, authorIDs []uint64, authorCh chan []uint64, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	var foundAuthorIDs []uint64
	for _, authorID := range authorIDs {
		author, err := b.authorRepo.FindByID(ctx, authorID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to find authors due %s`, err.Error()), err)
			return
		}
		foundAuthorIDs = append(foundAuthorIDs, author.ID)
	}
	authorCh <- foundAuthorIDs
}

// findGenres fetches genres by their IDs concurrently
func (b *bookUsecase) findGenres(ctx context.Context, genreIDs []uint64, genreCh chan []uint64, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	var foundGenreIDs []uint64
	for _, genreID := range genreIDs {
		genre, err := b.genreRepo.FindByID(ctx, genreID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to find genres due %s`, err.Error()), err)
			return
		}
		foundGenreIDs = append(foundGenreIDs, genre.ID)
	}
	genreCh <- foundGenreIDs
}

// createBook inserts a new book and sends the result through a channel
func (b *bookUsecase) createBook(ctx context.Context, request *requests.CreateBookRequest, bookCh chan *models.Book, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	bookImage := utils.RenameFileWithDateAndUUID(request.Image)
	err := utils.ResizeImage(request.Image, bookImage, 640, 1080)
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to resize image`, err)
		return
	}
	buckets, err := b.awsClient.ListBucket()
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to list buckets due %s`, err.Error()), err)
		return
	}
	imagePath, err := b.awsClient.UploadFile(s3manager.NewUploader(b.awsClient.SessionClient()), "books/"+bookImage, buckets.Buckets[0].String(), bookImage)
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to upload book image`, err)
		return
	}
	book, err := b.bookRepo.Save(ctx, &models.Book{
		Title:       request.Title,
		Description: request.Description,
		Status:      enums.Available,
		Image:       *imagePath,
	})
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create book due %s`, err.Error()), err)
		return
	}
	bookCh <- book
}

// createAuthorBookRelations creates relations between authors and the book
func (b *bookUsecase) createAuthorBookRelations(ctx context.Context, authorIDs []uint64, bookID uint64, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("author book relation insert")
	for _, authorID := range authorIDs {
		_, err := b.authorBookRepo.Save(ctx, &models.AuthorBook{
			AuthorID: authorID,
			BookID:   bookID,
		})
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create author book relation due %s`, err.Error()), err)
			return
		}
	}
}

// createBookGenreRelations creates relations between genres and the book
func (b *bookUsecase) createBookGenreRelations(ctx context.Context, genreIDs []uint64, bookID uint64, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, genreID := range genreIDs {
		_, err := b.bookGenreRepo.Save(ctx, &models.BookGenre{
			GenreID: genreID,
			BookID:  bookID,
		})
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, "failed to create book-genre relation", err)
			return
		}
	}
}

func (b *bookUsecase) GetAllBooks(ctx context.Context) (result []*responses.ListBooksResponse, customErr *apperror.CustomError) {
	result = []*responses.ListBooksResponse{}
	books, err := b.bookRepo.FindAll(ctx, 0, 0)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get books`, err)
	}
	for _, book := range books {
		result = append(result, &responses.ListBooksResponse{
			ID:    book.ID,
			Image: book.Image,
			Title: book.Title,
		})
	}
	return
}

func (b *bookUsecase) GetBookByID(ctx context.Context, id uint64) (result *responses.BookDetailResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.BookUsecase = &bookUsecase{}

func newBookUsecase(
	userRepo repositories_interfaces.UserRepository,
	bookRepo repositories_interfaces.BookRepository,
	genreRepo repositories_interfaces.GenreRepository,
	authorRepo repositories_interfaces.AuthorRepository,
	authorBookRepo repositories_interfaces.AuthorBookRepository,
	bookGenreRepo repositories_interfaces.BookGenreRepository) *bookUsecase {
	return &bookUsecase{
		userRepo:       userRepo,
		bookRepo:       bookRepo,
		genreRepo:      genreRepo,
		authorRepo:     authorRepo,
		authorBookRepo: authorBookRepo,
		bookGenreRepo:  bookGenreRepo,
	}
}
