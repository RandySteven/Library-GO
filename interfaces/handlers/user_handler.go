package handlers_interfaces

import (
	"net/http"
)

type UserHandler interface {
	GetUserDetail(w http.ResponseWriter, r *http.Request)
	GetListOfUsers(w http.ResponseWriter, r *http.Request)
}
