package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"net/http"
)

type UserHandler struct {
	usecase usecases_interfaces.UserUsecase
}

func (u *UserHandler) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) GetListOfUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.UserHandler = &UserHandler{}

func newUserHandler(usecase usecases_interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
