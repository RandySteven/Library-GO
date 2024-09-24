package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type BagUsecase interface {
	AddBookToBag(ctx context.Context, request *requests.BagRequest) (result *responses.AddBagResponse, customErr *apperror.CustomError)
	GetUserBag(ctx context.Context) (result *responses.GetAllBagsResponse, customErr *apperror.CustomError)
	DeleteBookFromBag(ctx context.Context, request *requests.BagRequest) (customErr *apperror.CustomError)
	EmptyBag(ctx context.Context) (customErr *apperror.CustomError)
}
