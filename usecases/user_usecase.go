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
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}

	result = &responses.UserDetailResponse{
		ID:             user.ID,
		Name:           user.Name,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		ProfilePicture: user.ProfilePicture,
		DoB:            user.DoB,
	}

	return result, nil
}

var _ usecases_interfaces.UserUsecase = &userUsecase{}

func newUserUsecase(userRepo repositories_interfaces.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
