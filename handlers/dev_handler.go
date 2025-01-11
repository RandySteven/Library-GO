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

type DevHandler struct {
	usecase usecases_interfaces.DevUsecase
}

func (d *DevHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.ResponseHandler(w, http.StatusOK, `success health check`, nil, nil, nil)
}

func (d *DevHandler) CheckMessageBroker(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `result`
	)
	result, err := d.usecase.MessageBrokerCheckerHealth(ctx)
	if err != nil {
		utils.ResponseHandler(w, http.StatusInternalServerError, `the message broker still issue`, nil, nil, err)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success health check`, &dataKey, result, nil)
}

func (d *DevHandler) CreateBucket(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CreateBucketRequest{}
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	err := d.usecase.CreateBucket(ctx, request.BucketName)
	if err != nil {
		utils.ResponseHandler(w, http.StatusInternalServerError, `internal server error`, nil, nil, err)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success create bucket`, nil, nil, nil)
}

func (d *DevHandler) GetListBuckets(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `buckets`
	)
	result, err := d.usecase.GetListBuckets(ctx)
	if err != nil {
		utils.ResponseHandler(w, http.StatusInternalServerError, `internal server error`, nil, nil, err)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get bucket`, &dataKey, result, nil)
}

var _ handlers_interfaces.DevHandler = &DevHandler{}

func newDevHandler(usecase usecases_interfaces.DevUsecase) *DevHandler {
	return &DevHandler{
		usecase: usecase,
	}
}
