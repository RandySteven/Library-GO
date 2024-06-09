package usecases

import (
	repositories_interfaces "github.com/RandySteven/Library-GO/user_service/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/user_service/interfaces/usecases"
)

type userUsecase struct {
	userRepo    repositories_interfaces.UserRepository
	profileRepo repositories_interfaces.UserProfileRepository
}

var _ usecases_interfaces.UserUsecase = &userUsecase{}
