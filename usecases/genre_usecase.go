package usecases

import (
	"context"
	"errors"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/redis/go-redis/v9"
	"sync"
)

type genreUsecase struct {
	genreRepo     repositories_interfaces.GenreRepository
	bookRepo      repositories_interfaces.BookRepository
	bookGenreRepo repositories_interfaces.BookGenreRepository
	ratingRepo    repositories_interfaces.RatingRepository
	genreCache    caches_interfaces.GenreCache
	userGenreRepo repositories_interfaces.UserGenreRepository
}

func (g *genreUsecase) GetGenreDetail(ctx context.Context, id uint64) (result *responses.GenreResponseDetail, customErr *apperror.CustomError) {
	var (
		bookResponses = []*responses.ListBooksResponse{}
	)
	genre, err := g.genreRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get genre id`, err)
	}

	bookGenres, err := g.bookGenreRepo.FindBookGenreByGenreID(ctx, genre.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get genre id`, err)
	}

	if len(bookGenres) != 0 {
		for _, book := range bookGenres {
			rating, _ := g.ratingRepo.FindRatingForBook(ctx, book.Book.ID)
			if rating == nil {
				rating = &models.Rating{
					Score: 0,
				}
			}
			bookResponses = append(bookResponses, &responses.ListBooksResponse{
				ID:        book.Book.ID,
				Image:     book.Book.Image,
				Title:     book.Book.Title,
				Rating:    rating.Score,
				Status:    book.Book.Status.ToString(),
				CreatedAt: book.Book.CreatedAt.Local(),
				UpdatedAt: book.Book.UpdatedAt.Local(),
				DeletedAt: book.Book.DeletedAt,
			})
		}
	}

	result = &responses.GenreResponseDetail{
		ID:        genre.ID,
		Genre:     genre.Genre,
		Books:     bookResponses,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}

	return result, nil
}

func (g *genreUsecase) AddGenre(ctx context.Context, request *requests.GenreRequest) (idHash string, customErr *apperror.CustomError) {
	genre, err := g.genreRepo.Save(ctx, &models.Genre{
		Genre: request.Genre,
	})
	if err != nil {
		return "", apperror.NewCustomError(apperror.ErrInternalServer, `failed to create genre`, err)
	}
	return utils.HashID(genre.ID), nil
}

func (g *genreUsecase) GetAllGenres(ctx context.Context) (result []*responses.ListGenresResponse, customErr *apperror.CustomError) {
	result, err := g.genreCache.GetMultiData(ctx)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			genres, err := g.genreRepo.FindAll(ctx, 0, 0)
			if err != nil {
				return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all genres`, err)
			}
			for _, genre := range genres {
				result = append(result, &responses.ListGenresResponse{
					ID:    genre.ID,
					Genre: genre.Genre,
				})
			}
			_ = g.genreCache.SetMultiData(ctx, result)
			return
		}
	}

	return result, nil
}

func (g *genreUsecase) ChooseFavoriteGenres(ctx context.Context, request *requests.ChooseFavoriteGenresRequest) (result []*responses.ListGenresResponse, customErr *apperror.CustomError) {
	var (
		wg          = sync.WaitGroup{}
		customErrCh = make(chan *apperror.CustomError)
		resultCh    = make(chan []*responses.ListGenresResponse)
	)

	userId := ctx.Value(enums.UserID).(uint64)

	genres, err := g.genreRepo.FindSelectedGenresByID(ctx, request.GenreIDs)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get genre id`, err)
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, genreId := range request.GenreIDs {
			_, err = g.userGenreRepo.Save(ctx, &models.UserGenre{
				UserID:  userId,
				GenreID: genreId,
			})
			if err != nil {
				customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to add genre favorite`, err)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for _, genre := range genres {
			result = append(result, &responses.ListGenresResponse{
				ID:    genre.ID,
				Genre: genre.Genre,
			})
		}
		resultCh <- result
	}()

	go func() {
		wg.Wait()
		close(resultCh)
		close(customErrCh)
	}()

	select {
	case <-customErrCh:
		return nil, <-customErrCh
	default:
		return <-resultCh, nil
	}
}

var _ usecases_interfaces.GenreUsecase = &genreUsecase{}

func newGenreUsecase(genreRepo repositories_interfaces.GenreRepository,
	bookRepo repositories_interfaces.BookRepository,
	bookGenreRepo repositories_interfaces.BookGenreRepository,
	ratingRepo repositories_interfaces.RatingRepository,
	genreCache caches_interfaces.GenreCache) *genreUsecase {
	return &genreUsecase{
		genreRepo:     genreRepo,
		bookRepo:      bookRepo,
		bookGenreRepo: bookGenreRepo,
		ratingRepo:    ratingRepo,
		genreCache:    genreCache,
	}
}
