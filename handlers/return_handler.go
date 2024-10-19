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

type ReturnHandler struct {
	usecase usecases_interfaces.ReturnUsecase
}

func (r2 *ReturnHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.ReturnRequest{}
		dataKey = `result`
	)
	if err := utils.BindRequest(r, request); err != nil {
		return
	}
	result, customErr := r2.usecase.ReturnBook(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to return book`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success return book`, &dataKey, result, nil)
}

var _ handlers_interfaces.ReturnHandler = &ReturnHandler{}

func newReturnHandler(usecase usecases_interfaces.ReturnUsecase) *ReturnHandler {
	return &ReturnHandler{
		usecase: usecase,
	}
}
