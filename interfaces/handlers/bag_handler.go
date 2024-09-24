package handlers_interfaces

import "net/http"

type BagHandler interface {
	AddBookToBag(w http.ResponseWriter, r *http.Request)
	GetUserBag(w http.ResponseWriter, r *http.Request)
	DeleteBookFromBag(w http.ResponseWriter, r *http.Request)
}
