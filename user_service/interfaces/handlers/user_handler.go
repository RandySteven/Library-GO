package handlers_interfaces

import "net/http"

type UserHandler interface {
	LoginUser(w http.ResponseWriter, r *http.Request)
	RegisterUser(w http.ResponseWriter, r *http.Request)
	UserProfile(w http.ResponseWriter, r *http.Request)
	UserDetail(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	RemoveUser(w http.ResponseWriter, r *http.Request)
}
