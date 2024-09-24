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

type BagHandler struct {
	usecase usecases_interfaces.BagUsecase
}

func (b *BagHandler) AddBookToBag(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.BagRequest{}
		dataKey = `book`
	)
	if err := utils.BindRequest(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := b.usecase.AddBookToBag(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success add book to bag`, &dataKey, result, nil)
}

func (b *BagHandler) GetUserBag(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `bag`
	)
	result, customErr := b.usecase.GetUserBag(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get user bag`, &dataKey, result, nil)
}

func (b *BagHandler) DeleteBookFromBag(w http.ResponseWriter, r *http.Request) {
}

var _ handlers_interfaces.BagHandler = &BagHandler{}

func newBagHandler(usecase usecases_interfaces.BagUsecase) *BagHandler {
	return &BagHandler{
		usecase: usecase,
	}
}
