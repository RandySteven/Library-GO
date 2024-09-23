package handlers_interfaces

import "net/http"

type GenreHandler interface {
	AddNewGenre(w http.ResponseWriter, r *http.Request)
	GetAllGenres(w http.ResponseWriter, r *http.Request)
}
