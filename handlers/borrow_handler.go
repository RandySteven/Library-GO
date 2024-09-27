package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"net/http"
)

type BorrowHandler struct {
	usecase usecases_interfaces.BorrowUsecase
}

func (b *BorrowHandler) BorrowCheckout(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.BorrowHandler = &BorrowHandler{}

func newBorrowHandler(usecase usecases_interfaces.BorrowUsecase) *BorrowHandler {
	return &BorrowHandler{
		usecase: usecase,
	}
}
