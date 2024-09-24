package handlers

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/enums"
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookHandler struct {
	usecase usecases_interfaces.BookUsecase
}

func (b *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), "rID", rID)
		dataKey = `book`
	)

	request := &requests.CreateBookRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Genres:      utils.SeparateStringIntoUint64Arr(r.FormValue("genres"), ","),
		Authors:     utils.SeparateStringIntoUint64Arr(r.FormValue("authors"), ","),
	}

	imageFile, fileHeader, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer imageFile.Close()
	request.Image = imageFile

	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	//fileHeader := r.MultipartForm.File["image"][0]

	result, customErr := b.usecase.AddNewBook(ctx, request, fileHeader)
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
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `book`
		vars    = mux.Vars(r)
		idStr   = vars[`id`]
	)
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := b.usecase.GetBookByID(ctx, idUint64)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get book`, &dataKey, result, nil)
}

var _ handlers_interfaces.BookHandler = &BookHandler{}

func newBookHandler(usecase usecases_interfaces.BookUsecase) *BookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}
