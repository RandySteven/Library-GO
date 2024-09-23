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
	genreRepo repositories_interfaces.GenreRepository
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

func newGenreUsecase(genreRepo repositories_interfaces.GenreRepository) *genreUsecase {
	return &genreUsecase{genreRepo: genreRepo}
}
