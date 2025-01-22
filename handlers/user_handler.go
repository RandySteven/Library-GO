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

type UserHandler struct {
	usecase usecases_interfaces.UserUsecase
}

func (u *UserHandler) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		params  = mux.Vars(r)
		dataKey = `user`
	)
	idStr := params[`id`]
	idUint, _ := strconv.Atoi(idStr)
	result, customErr := u.usecase.GetUserDetail(ctx, uint64(idUint))
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get users`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get user`, &dataKey, result, nil)
}

func (u *UserHandler) GetListOfUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.UserHandler = &UserHandler{}

func newUserHandler(usecase usecases_interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
