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

type CommentHandler struct {
	usecase usecases_interfaces.CommentUsecase
}

func (c *CommentHandler) Comment(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.AddCommentRequest{}
		dataKey = `result`
	)
	if err := utils.BindRequest(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `invalid request`, nil, nil, err)
		return
	}
	result, customErr := c.usecase.AddComment(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to add comment`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success add comment`, &dataKey, result, nil)
}

func (c *CommentHandler) Reply(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentHandler) GetBookComment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.CommentHandler = &CommentHandler{}

func newCommentHandler(usecase usecases_interfaces.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		usecase: usecase,
	}
}
