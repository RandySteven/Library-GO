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

type RatingHandler struct {
	usecases usecases_interfaces.RatingUsecase
}

func (r2 *RatingHandler) SubmitRating(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.RatingRequest{}
		dataKey = `rating`
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := r2.usecases.SubmitBookRating(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success register user`, &dataKey, result, nil)
}

var _ handlers_interfaces.RatingHandler = &RatingHandler{}

func newRatingHandler(usecases usecases_interfaces.RatingUsecase) *RatingHandler {
	return &RatingHandler{
		usecases: usecases,
	}
}
