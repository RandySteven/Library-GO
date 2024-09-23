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
}
