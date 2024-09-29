package handlers_interfaces

import "net/http"

type BookHandler interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetBookByID(w http.ResponseWriter, r *http.Request)
	SearchBooks(w http.ResponseWriter, r *http.Request)
}
