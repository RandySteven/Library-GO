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

type OnboardingHandler struct {
	onboardingUsecase usecases_interfaces.OnboardingUsecase
}

func (o *OnboardingHandler) GetLoginUser(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `user`
	)
	result, customErr := o.onboardingUsecase.GetLoginUser(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success register user`, &dataKey, result, nil)
}

func (o *OnboardingHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.UserRegisterRequest{}
		dataKey = `user`
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := o.onboardingUsecase.RegisterUser(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success register user`, &dataKey, result, nil)
}

func (o *OnboardingHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.UserLoginRequest{}
		dataKey = `user`
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `bad request`, nil, nil, err)
		return
	}
	result, customErr := o.onboardingUsecase.LoginUser(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal server error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success login user`, &dataKey, result, nil)
}

func (o *OnboardingHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
}

var _ handlers_interfaces.OnboardingHandler = &OnboardingHandler{}

func newOnboardingHandler(onboardingUsecase usecases_interfaces.OnboardingUsecase) *OnboardingHandler {
	return &OnboardingHandler{
		onboardingUsecase: onboardingUsecase,
	}
}
