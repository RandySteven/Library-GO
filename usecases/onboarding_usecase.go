package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type onboardingUsecase struct {
	userRepo repositories_interfaces.UserRepository
}

func (o onboardingUsecase) RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (result *responses.UserRegisterResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (o onboardingUsecase) LoginUser(ctx context.Context, request *requests.UserLoginRequest) (result *responses.UserLoginResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (o onboardingUsecase) VerifyToken(ctx context.Context, token string) (customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

var _ usecases_interfaces.OnboardingUsecase = &onboardingUsecase{}

func newOnboardingUsecase(userRepo repositories_interfaces.UserRepository) *onboardingUsecase {
	return &onboardingUsecase{
		userRepo: userRepo,
	}
}
