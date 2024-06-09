package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/user_service/entities/payloads/requests"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, request *requests.UserRegisterRequest)
	LoginUser(ctx context.Context, request *requests.UserLoginRequest)
}
