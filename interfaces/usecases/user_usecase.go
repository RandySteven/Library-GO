package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type UserUsecase interface {
	GetAllUsers(ctx context.Context) (result []*responses.UserListResponse, customErr *apperror.CustomError)
	GetUserDetail(ctx context.Context, id uint64) (result *responses.UserDetailResponse, customErr *apperror.CustomError)
	GetLoginUser(ctx context.Context) (result *responses.LoginUserResponse, customErr *apperror.CustomError)
}
