package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type GenreUsecase interface {
	AddGenre(ctx context.Context, request *requests.GenreRequest) (idHash string, customErr *apperror.CustomError)
	GetAllGenres(ctx context.Context) (result []*responses.ListGenresResponse, customErr *apperror.CustomError)
	GetGenreDetail(ctx context.Context, id uint64) (result *responses.GenreResponseDetail, customErr *apperror.CustomError)
	ChooseFavoriteGenres(ctx context.Context, request *requests.ChooseFavoriteGenresRequest) (result []*responses.ListGenresResponse, customErr *apperror.CustomError)
}
