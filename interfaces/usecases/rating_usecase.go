package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type RatingUsecase interface {
	SubmitBookRating(ctx context.Context, request *requests.RatingRequest) (result *responses.RatingResponse, customErr *apperror.CustomError)
}
