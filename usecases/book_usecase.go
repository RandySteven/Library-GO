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
	algolia_client "github.com/RandySteven/Library-GO/pkg/algolia"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"sync"
)

type bookUsecase struct {
	awsClient      *aws_client.AWSClient
	algoClient     *algolia_client.AlgoliaAPISearchClient
	userRepo       repositories_interfaces.UserRepository
	bookRepo       repositories_interfaces.BookRepository
	genreRepo      repositories_interfaces.GenreRepository
	authorRepo     repositories_interfaces.AuthorRepository
	authorBookRepo repositories_interfaces.AuthorBookRepository
	bookGenreRepo  repositories_interfaces.BookGenreRepository
	ratingRepo     repositories_interfaces.RatingRepository
}

func (b *bookUsecase) refreshTx(ctx context.Context) {
	tx := b.bookRepo.GetTx(ctx)
	b.userRepo.SetTx(tx)
	b.genreRepo.SetTx(tx)
	b.authorRepo.SetTx(tx)
	b.authorBookRepo.SetTx(tx)
	b.bookGenreRepo.SetTx(tx)
	b.ratingRepo.SetTx(tx)
}

func (b *bookUsecase) AddNewBook(ctx context.Context, request *requests.CreateBookRequest, fileHeader *multipart.FileHeader) (result *responses.CreateBookResponse, customErr *apperror.CustomError) {
	var (
		wg       sync.WaitGroup
		wg2      sync.WaitGroup
		errCh    = make(chan *apperror.CustomError, 1) // Buffered channels to prevent blocking
		errCh2   = make(chan *apperror.CustomError, 1)
		bookCh   = make(chan *models.Book, 1)     // Buffered channel
		authorCh = make(chan []*models.Author, 1) // Buffered channel
		genreCh  = make(chan []*models.Genre, 1)  // Buffered channel
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
		b.bookRepo.SetTx(nil)
	}()
	b.refreshTx(ctx)

	// 2. Search authors, genres, and insert book concurrently
	wg.Add(3)
	go b.findAuthors(ctx, request.Authors, authorCh, errCh, &wg)
	go b.findGenres(ctx, request.Genres, genreCh, errCh, &wg)
	go b.createBook(ctx, request, bookCh, errCh, &wg, fileHeader)

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errCh)
		close(authorCh)
		close(genreCh)
		close(bookCh)
	}()

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
	hashedID := utils.HashID(book.ID)
	_, err = b.algoClient.SaveObject(enums.BooksIndex, map[string]any{
		"objectID":    book.ID,
		"title":       book.Title,
		"description": book.Description,
		"image":       book.Image,
		"genres":      genreIDs,
		"authors":     authorIDs,
		"createdAt":   book.CreatedAt.Local(),
		"updatedAt":   book.UpdatedAt.Local(),
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to save book", err)
	}

	result = &responses.CreateBookResponse{
		ID: hashedID,
	}
	return result, nil
}

// findAuthors fetches authors by their IDs concurrently
func (b *bookUsecase) findAuthors(ctx context.Context, authorIDs []uint64, authorCh chan []*models.Author, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	var foundAuthorIDs []*models.Author
	for _, authorID := range authorIDs {
		author, err := b.authorRepo.FindByID(ctx, authorID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to find authors due %s`, err.Error()), err)
			return
		}
		foundAuthorIDs = append(foundAuthorIDs, author)
	}
	authorCh <- foundAuthorIDs
}

// findGenres fetches genres by their IDs concurrently
func (b *bookUsecase) findGenres(ctx context.Context, genreIDs []uint64, genreCh chan []*models.Genre, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	var foundGenreIDs []*models.Genre
	for _, genreID := range genreIDs {
		genre, err := b.genreRepo.FindByID(ctx, genreID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to find genres due %s`, err.Error()), err)
			return
		}
		foundGenreIDs = append(foundGenreIDs, genre)
	}
	genreCh <- foundGenreIDs
}

// createBook inserts a new book and sends the result through a channel
func (b *bookUsecase) createBook(ctx context.Context, request *requests.CreateBookRequest, bookCh chan *models.Book, errCh chan *apperror.CustomError, wg *sync.WaitGroup, fileHeader *multipart.FileHeader) {
	defer wg.Done()

	tempFile, err := ioutil.TempFile("./temp-images", "upload-*.png")
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, "temp-images required", err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(request.Image)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	if fileHeader == nil {
		errCh <- apperror.NewCustomError(apperror.ErrBadRequest, "image is required", nil)
		return
	}

	imageFile, err := fileHeader.Open()
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, "failed to open image file", err)
		return
	}
	defer imageFile.Close()

	renamedImage := utils.RenameFileWithDateAndUUID(tempFile.Name()[len(`./temp-images/`):])

	buckets, err := b.awsClient.ListBucket()
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf("failed to list buckets: %s", err), err)
		return
	}

	imagePath, err := b.awsClient.UploadFile(s3manager.NewUploader(b.awsClient.SessionClient()), tempFile.Name(), *buckets.Buckets[0].Name, "books/"+renamedImage)
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, "failed to upload book image", err)
		return
	}
	_ = os.Remove(tempFile.Name())

	// Save the book information to the repository
	book, err := b.bookRepo.Save(ctx, &models.Book{
		Title:       request.Title,
		Description: request.Description,
		Status:      enums.Available,
		Image:       *imagePath,
	})
	if err != nil {
		errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf("failed to create book: %s", err), err)
		return
	}

	// Send the created book back through the channel
	bookCh <- book
}

// createAuthorBookRelations creates relations between authors and the book
func (b *bookUsecase) createAuthorBookRelations(ctx context.Context, authorIDs []*models.Author, bookID uint64, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("author book relation insert")
	for _, authorID := range authorIDs {
		_, err := b.authorBookRepo.Save(ctx, &models.AuthorBook{
			AuthorID: authorID.ID,
			BookID:   bookID,
		})
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, fmt.Sprintf(`failed to create author book relation due %s`, err.Error()), err)
			return
		}
	}
}

// createBookGenreRelations creates relations between genres and the book
func (b *bookUsecase) createBookGenreRelations(ctx context.Context, genreIDs []*models.Genre, bookID uint64, errCh chan *apperror.CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, genreID := range genreIDs {
		_, err := b.bookGenreRepo.Save(ctx, &models.BookGenre{
			GenreID: genreID.ID,
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
	return result, nil
}

func (b *bookUsecase) GetBookByID(ctx context.Context, id uint64) (result *responses.BookDetailResponse, customErr *apperror.CustomError) {
	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
		genresCh    = make(chan []string)
		authorsCh   = make(chan []string)
		ratingCh    = make(chan *models.Rating)
	)
	book, err := b.bookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get book by id`, err)
	}

	result = &responses.BookDetailResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Image:       book.Image,
		Status:      book.Status.ToString(),
		CreatedAt:   book.CreatedAt.Local(),
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		var authorNames []string
		authorBooks, err := b.authorBookRepo.FindAuthorBookByBookID(ctx, book.ID)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get author book by id`, err)
			return
		}
		var authorIDs []uint64
		for _, authorBook := range authorBooks {
			authorIDs = append(authorIDs, authorBook.AuthorID)
		}
		authors, err := b.authorRepo.FindSelectedAuthorsByID(ctx, authorIDs)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get author book by id`, err)
			return
		}
		for _, author := range authors {
			authorNames = append(authorNames, author.Name)
		}
		authorsCh <- authorNames
	}()

	go func() {
		defer wg.Done()
		genreNames := []string{}
		bookGenres, err := b.bookGenreRepo.FindBookGenreByBookID(ctx, book.ID)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get book genre by id`, err)
			return
		}
		genreIDs := []uint64{}
		for _, bookGenre := range bookGenres {
			genreIDs = append(genreIDs, bookGenre.ID)
		}

		genres, err := b.genreRepo.FindSelectedGenresByID(ctx, genreIDs)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get book genre by id`, err)
			return
		}
		for _, genre := range genres {
			genreNames = append(genreNames, genre.Genre)
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
	//searchResults, err := b.algoClient.Search(enums.BooksIndex, request.Keyword)
	//if err != nil {
	//	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to search books", err)
	//}
	//
	//result = []*responses.ListBooksResponse{}
	//
	//for _, res := range searchResults.Results {
	//	hits := res.SearchResponse.Hits
	//	for _, hit := range hits {
	//		mapHit := *hit.HighlightResult
	//		convertId, _ := strconv.Atoi(mapHit["objectID"].HighlightResultOption.Value)
	//		bookObj := &responses.ListBooksResponse{
	//			ID:    uint64(convertId),
	//			Title: mapHit["title"].HighlightResultOption.Value,
	//			Image: mapHit["image"].HighlightResultOption.Value,
	//		}
	//		result = append(result, bookObj)
	//	}
	//}

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
	ratingRepo repositories_interfaces.RatingRepository,
	awsClient *aws_client.AWSClient,
	algoClient *algolia_client.AlgoliaAPISearchClient) *bookUsecase {
	return &bookUsecase{
		userRepo:       userRepo,
		bookRepo:       bookRepo,
		genreRepo:      genreRepo,
		authorRepo:     authorRepo,
		authorBookRepo: authorBookRepo,
		bookGenreRepo:  bookGenreRepo,
		ratingRepo:     ratingRepo,
		awsClient:      awsClient,
		algoClient:     algoClient,
	}
}
