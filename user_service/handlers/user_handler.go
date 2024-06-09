package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/user_service/interfaces/handlers"
	"net/http"
)

type UserHandler struct{}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) UserProfile(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) UserDetail(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.UserHandler = &UserHandler{}
