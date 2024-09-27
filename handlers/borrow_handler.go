package handlers

import (
	"context"
	"github.com/RandySteven/Library-GO/enums"
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BorrowHandler struct {
	usecase usecases_interfaces.BorrowUsecase
}

func (b *BorrowHandler) GetBorrowDetail(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `result`
	)
	param := mux.Vars(r)
	idStr, err := strconv.Atoi(param[`id`])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, customErr := b.usecase.GetBorrowDetail(ctx, uint64(idStr))
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to borrow book`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to borrow book`, &dataKey, result, nil)
}

func (b *BorrowHandler) BorrowCheckout(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `result`
	)
	result, customErr := b.usecase.BorrowTransaction(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to borrow book`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to borrow book`, &dataKey, result, nil)
}

var _ handlers_interfaces.BorrowHandler = &BorrowHandler{}

func newBorrowHandler(usecase usecases_interfaces.BorrowUsecase) *BorrowHandler {
	return &BorrowHandler{
		usecase: usecase,
	}
}
