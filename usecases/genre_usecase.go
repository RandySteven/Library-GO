package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type genreUsecase struct {
	genreRepo     repositories_interfaces.GenreRepository
	bookRepo      repositories_interfaces.BookRepository
	bookGenreRepo repositories_interfaces.BookGenreRepository
}

func (g *genreUsecase) GetGenreDetail(ctx context.Context, id uint64) (result *responses.GenreResponseDetail, customErr *apperror.CustomError) {
	var (
		bookIDs       = []uint64{}
		bookResponses = []*responses.ListBooksResponse{}
	)
	genre, err := g.genreRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get genre id`, err)
	}

	bookGenres, err := g.bookGenreRepo.FindBookGenreByBookID(ctx, genre.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get genre id`, err)
	}

	for _, bookGenre := range bookGenres {
		bookIDs = append(bookIDs, bookGenre.ID)
	}

	if len(bookIDs) != 0 {
		books, err := g.bookRepo.FindSelectedBooksId(ctx, bookIDs)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get genre id`, err)
		}
		for _, book := range books {
			bookResponses = append(bookResponses, &responses.ListBooksResponse{
				ID:        book.ID,
				Image:     book.Image,
				Title:     book.Title,
				Status:    book.Status.ToString(),
				CreatedAt: book.CreatedAt.Local(),
				UpdatedAt: book.UpdatedAt.Local(),
				DeletedAt: book.DeletedAt,
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
	//TODO implement me
	panic("implement me")
}

func (g *genreUsecase) GetAllGenres(ctx context.Context) (result []*responses.ListGenresResponse, customErr *apperror.CustomError) {
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
	return result, nil
}

var _ usecases_interfaces.GenreUsecase = &genreUsecase{}

func newGenreUsecase(genreRepo repositories_interfaces.GenreRepository,
	bookRepo repositories_interfaces.BookRepository,
	bookGenreRepo repositories_interfaces.BookGenreRepository) *genreUsecase {
	return &genreUsecase{genreRepo: genreRepo, bookRepo: bookRepo, bookGenreRepo: bookGenreRepo}
}
