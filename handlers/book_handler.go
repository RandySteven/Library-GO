package handlers

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/enums"
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"net/http"
)

type BookHandler struct {
	usecase usecases_interfaces.BookUsecase
}

func (b *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), "rID", rID)
		request = &requests.CreateBookRequest{}
		dataKey = `book`
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := b.usecase.AddNewBook(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success create book`, &dataKey, result, nil)
}

func (b *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `books`
	)
	result, customErr := b.usecase.GetAllBooks(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get books`, &dataKey, result, nil)
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
