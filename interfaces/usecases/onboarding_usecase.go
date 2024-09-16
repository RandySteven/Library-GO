package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type OnboardingUsecase interface {
	RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (result *responses.UserRegisterResponse, customErr *apperror.CustomError)
	LoginUser(ctx context.Context, request *requests.UserLoginRequest) (result *responses.UserLoginResponse, customErr *apperror.CustomError)
}
