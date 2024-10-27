package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type userUsecase struct {
	userRepo repositories_interfaces.UserRepository
}

func (u *userUsecase) GetAllUsers(ctx context.Context) (result []*responses.UserListResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (u *userUsecase) GetUserDetail(ctx context.Context, id uint64) (result *responses.UserDetailResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.UserUsecase = &userUsecase{}

func newUserUsecase(userRepo repositories_interfaces.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
