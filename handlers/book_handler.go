package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"net/http"
)

type BookHandler struct {
	usecase usecases_interfaces.BookUsecase
}

func (b *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.BookHandler = &BookHandler{}

func newBookHandler(usecase usecases_interfaces.BookUsecase) *BookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}
