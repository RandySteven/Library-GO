package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	algolia_client "github.com/RandySteven/Library-GO/pkg/algolia"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"mime/multipart"
	"strconv"
	"sync"
)

type bookUsecase struct {
	awsClient        aws_client.AWS
	algoClient       algolia_client.AlgoliaAPISearch
	userRepo         repositories_interfaces.UserRepository
	bookRepo         repositories_interfaces.BookRepository
	genreRepo        repositories_interfaces.GenreRepository
	authorRepo       repositories_interfaces.AuthorRepository
	authorBookRepo   repositories_interfaces.AuthorBookRepository
	bookGenreRepo    repositories_interfaces.BookGenreRepository
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDeatilRepo repositories_interfaces.BorrowDetailRepository
	ratingRepo       repositories_interfaces.RatingRepository
	cache            caches_interfaces.BookCache
	pubsub           rabbitmqs_client.PubSub
	transaction      repositories_interfaces.Transaction
}

func (b *bookUsecase) AddNewBook(ctx context.Context, request *requests.CreateBookRequest, fileHeader *multipart.FileHeader) (result *responses.CreateBookResponse, customErr *apperror.CustomError) {
	var (
		wg            sync.WaitGroup
		errCh         = make(chan *apperror.CustomError, 1)
		goroutineDone = make(chan struct{})
		book          = &models.Book{}
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	customErr = b.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {
		imagePath, err := b.awsClient.UploadImageFile(ctx, request.Image, enums.BooksPath, fileHeader, 600, 900)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, "failed to upload book image", err)
		}

		book, err = b.bookRepo.Save(ctx, &models.Book{
			Title:       request.Title,
			Description: request.Description,
			Status:      enums.Available,
			Image:       *imagePath,
		})
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf("failed to create book: %s", err), err)
		}

		wg.Add(2)

		go func() {
			defer wg.Done()
			if len(request.Authors) == 1 {
				_, err = b.authorBookRepo.Save(ctx, &models.AuthorBook{
					AuthorID: request.Authors[0],
					BookID:   book.ID,
				})
				if err != nil {
					select {
					case errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create author book relation due %s`, err.Error()), err):
						cancel()
					}
					return
				}
				return
			}
			for _, authorID := range request.Authors {
				_, err = b.authorBookRepo.Save(ctx, &models.AuthorBook{
					AuthorID: authorID,
					BookID:   book.ID,
				})
				if err != nil {
					log.Println("error ", err)
					select {
					case errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create author book relation due %s`, err.Error()), err):
						cancel()
					}
					return
				}
			}
		}()

		go func() {
			defer wg.Done()
			if len(request.Genres) == 1 {
				_, err = b.bookGenreRepo.Save(ctx, &models.BookGenre{
					GenreID: request.Genres[0],
					BookID:  book.ID,
				})
				if err != nil {
					select {
					case errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create book genre relation due %s`, err.Error()), err):
						cancel()
					}
					return
				}
				return
			}
			for _, genreID := range request.Genres {
				_, err = b.bookGenreRepo.Save(ctx, &models.BookGenre{
					GenreID: genreID,
					BookID:  book.ID,
				})
				if err != nil {
					select {
					case errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create book genre relation due %s`, err.Error()), err):
						cancel()
					}
					return
				}
			}
		}()

		go func() {
			wg.Wait()
			close(errCh)
			close(goroutineDone)
		}()

		select {
		case <-goroutineDone:
			var firstErr *apperror.CustomError
			for err := range errCh {
				if firstErr == nil {
					firstErr = err
				}
			}
			if firstErr != nil {
				return firstErr
			}
		case <-ctx.Done():
			return apperror.NewCustomError(apperror.ErrInternalServer, "context cancelled", ctx.Err())
		}
		return nil
	})
	if customErr != nil {
		return nil, customErr
	}

	_ = b.pubsub.Send(ctx, "book_exchange", "book-send-message", book)
	result = &responses.CreateBookResponse{
		ID: uuid.NewString(),
	}
	return result, nil
}

func (b *bookUsecase) GetAllBooks(ctx context.Context, request *requests.PaginationRequest) (result []*responses.ListBooksResponse, customErr *apperror.CustomError) {
	result, err := b.cache.GetBookPage(ctx, request.Page)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `redis issue : `, err)
		}
	}
	if result != nil {
		return result, nil
	}
	books, err := b.bookRepo.FindAll(ctx, request.Page, request.Limit)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get books`, err)
	}
	for _, book := range books {
		rating, err := b.ratingRepo.FindRatingForBook(ctx, book.ID)
		if err != nil {
			rating = &models.Rating{}
			rating.Score = 0
		}
		result = append(result, &responses.ListBooksResponse{
			ID:        book.ID,
			Image:     book.Image,
			Title:     book.Title,
			Rating:    rating.Score,
			Status:    book.Status.ToString(),
			CreatedAt: book.CreatedAt.Local(),
			UpdatedAt: book.UpdatedAt.Local(),
			DeletedAt: book.DeletedAt,
		})
	}
	err = b.cache.SetBookPage(ctx, request.Page, result)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert redis`, err)
	}
	return result, nil
}

func (b *bookUsecase) GetBookByID(ctx context.Context, id uint64) (result *responses.BookDetailResponse, customErr *apperror.CustomError) {
	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
		genresCh    = make(chan []*responses.GenreBookResponse)
		authorsCh   = make(chan []*responses.AuthorBookResponse)
		ratingCh    = make(chan *models.Rating)
	)
	book, err := b.bookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `failed to get book by id`, err)
	}

	result = &responses.BookDetailResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Image:       book.Image,
		PDFFile:     book.PDFFile,
		Status:      book.Status.ToString(),
		CreatedAt:   book.CreatedAt.Local(),
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		var authorNames = []*responses.AuthorBookResponse{}
		authorBooks, err := b.authorBookRepo.FindAuthorBookByBookID(ctx, book.ID)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get author book by id`, err)
			return
		}
		for _, author := range authorBooks {
			authorNames = append(authorNames, &responses.AuthorBookResponse{
				ID:   author.Author.ID,
				Name: author.Author.Name,
			})
		}
		authorsCh <- authorNames
	}()

	go func() {
		defer wg.Done()
		genreNames := []*responses.GenreBookResponse{}
		bookGenres, err := b.bookGenreRepo.FindBookGenreByBookID(ctx, book.ID)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get book genre by id`, err)
			return
		}
		for _, genre := range bookGenres {
			genreNames = append(genreNames, &responses.GenreBookResponse{
				ID:    genre.Genre.ID,
				Genre: genre.Genre.Genre,
			})
		}
		genresCh <- genreNames
	}()

	go func() {
		defer wg.Done()
		rating, err := b.ratingRepo.FindRatingForBook(ctx, book.ID)
		if err != nil {
			rating = &models.Rating{}
			rating.Score = 0
			ratingCh <- rating
			return
		}
		ratingCh <- rating
	}()

	go func() {
		wg.Wait()
		close(customErrCh)
		close(genresCh)
		close(authorsCh)
		close(ratingCh)
	}()

	select {
	case customErr = <-customErrCh:
		return nil, customErr
	default:
		result.Authors = <-authorsCh
		result.Genres = <-genresCh
		rating := <-ratingCh
		result.Rating = rating.Score
		return result, nil
	}
}

func (b *bookUsecase) SearchBook(ctx context.Context, request *requests.SearchBookRequest) (result []*responses.ListBooksResponse, customErr *apperror.CustomError) {
	searchResults, err := b.algoClient.Search(enums.BooksIndex, request.Keyword)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to search books", err)
	}

	result = []*responses.ListBooksResponse{}

	for _, res := range searchResults.Results {
		hits := res.SearchResponse.Hits
		for _, hit := range hits {
			mapHit := *hit.HighlightResult
			convertId, _ := strconv.Atoi(mapHit["objectID"].HighlightResultOption.Value)
			bookObj := &responses.ListBooksResponse{
				ID:    uint64(convertId),
				Title: mapHit["title"].HighlightResultOption.Value,
				Image: mapHit["image"].HighlightResultOption.Value,
			}
			result = append(result, bookObj)
		}
	}

	return result, nil
}

func (b *bookUsecase) BookBorrowTracker(ctx context.Context, id uint64) (result []*responses.BookBorrowHistoryResponse, customErr *apperror.CustomError) {
	borrowDetails, err := b.borrowDeatilRepo.FindByBorrowID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow details`, err)
	}

	for _, borrowDetail := range borrowDetails {
		borrow, err := b.borrowRepo.FindByID(ctx, borrowDetail.BorrowID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow`, err)
		}
		result = append(result, &responses.BookBorrowHistoryResponse{
			ID: borrowDetail.BorrowID,
			User: struct {
				ID   uint64 `json:"id"`
				Name string `json:"name"`
			}{ID: borrow.User.ID, Name: borrow.User.Name},
			BorrowDate: borrowDetail.BorrowedDate.Local(),
		})
	}

	return result, nil
}

var _ usecases_interfaces.BookUsecase = &bookUsecase{}

func newBookUsecase(
	userRepo repositories_interfaces.UserRepository,
	bookRepo repositories_interfaces.BookRepository,
	genreRepo repositories_interfaces.GenreRepository,
	authorRepo repositories_interfaces.AuthorRepository,
	authorBookRepo repositories_interfaces.AuthorBookRepository,
	bookGenreRepo repositories_interfaces.BookGenreRepository,
	borrowRepo repositories_interfaces.BorrowRepository,
	borrowDeatilRepo repositories_interfaces.BorrowDetailRepository,
	ratingRepo repositories_interfaces.RatingRepository,
	awsClient aws_client.AWS,
	algoClient algolia_client.AlgoliaAPISearch,
	cache caches_interfaces.BookCache,
	pubsub rabbitmqs_client.PubSub,
	transaction repositories_interfaces.Transaction) *bookUsecase {
	return &bookUsecase{
		userRepo:         userRepo,
		bookRepo:         bookRepo,
		genreRepo:        genreRepo,
		authorRepo:       authorRepo,
		authorBookRepo:   authorBookRepo,
		bookGenreRepo:    bookGenreRepo,
		borrowRepo:       borrowRepo,
		borrowDeatilRepo: borrowDeatilRepo,
		ratingRepo:       ratingRepo,
		awsClient:        awsClient,
		algoClient:       algoClient,
		cache:            cache,
		pubsub:           pubsub,
		transaction:      transaction,
	}
}
