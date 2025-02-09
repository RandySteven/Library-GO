package handlers

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type BookHandler struct {
	usecase usecases_interfaces.BookUsecase
}

func (b *BookHandler) BookBorrowHistoryTracker(w http.Request, r *http.Request) {

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
		dataKey = `result`
		request = &requests.PaginationRequest{}
	)
	request = paginationRequest(r.URL.Query())

	result, customErr := b.usecase.GetAllBooks(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	pageInt := request.Page
	limitInt := request.Limit

	next, prev := prevNext(pageInt, limitInt, len(result))

	response := &responses.PaginationListBookResponse{
		Books: result,
		Next:  next,
		Prev:  prev,
	}
	utils.ResponseHandler(w, http.StatusOK, `success get books`, &dataKey, response, nil)
}

func (b *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `book`
		vars    = mux.Vars(r)
		idStr   = vars[`id`]
		//id      = r.PathValue("id")
	)
	log.Println(idStr)
	idUint64, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := b.usecase.GetBookByID(ctx, uint64(idUint64))
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get book`, &dataKey, result, nil)
}

func (b *BookHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `book`
		request = &requests.SearchBookRequest{}
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := b.usecase.SearchBook(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success search books`, &dataKey, result, nil)
}

func paginationRequest(query url.Values) (request *requests.PaginationRequest) {
	request = &requests.PaginationRequest{}
	page := query.Get(`page`)
	limit := query.Get(`limit`)
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	request.Page = uint64(pageInt)
	request.Limit = uint64(limitInt)
	return request
}

func prevNext(pageInt, limitInt uint64, currResultLen int) (next, prev string) {
	prev = fmt.Sprintf("http://%s:%s/books?page=%d&limit=%d", os.Getenv("HOST"), os.Getenv("PORT"), pageInt-1, limitInt)
	if pageInt == 1 {
		prev = ""
	}
	next = fmt.Sprintf("http://%s:%s/books?page=%d&limit=%d", os.Getenv("HOST"), os.Getenv("PORT"), pageInt+1, limitInt)
	if currResultLen < int(limitInt) {
		next = ""
	}
	return next, prev
}

var _ handlers_interfaces.BookHandler = &BookHandler{}

func newBookHandler(usecase usecases_interfaces.BookUsecase) *BookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}
