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

type (
	RoomChatHandler struct {
		usecase usecases_interfaces.ChatUsecase
	}

	SendChatHandler struct {
		usecase usecases_interfaces.ChatUsecase
	}
)

func (s *SendChatHandler) SendChat(w http.ResponseWriter, r *http.Request) {
}

func (r2 *RoomChatHandler) CreateRoomChat(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CreateChatRoom{}
		dataKey = `room`
	)
	if err := utils.BindRequest(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `failed to get request`, nil, nil, err)
		return
	}
	result, customErr := r2.usecase.CreateRoomChat(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get request`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to create room chat`, &dataKey, result, nil)
}

func (r2 *RoomChatHandler) GetAllRoomsChat(w http.Response, r *http.Request) {
}

func (r2 *RoomChatHandler) InvitePeopleToRoom(w http.Response, r *http.Request) {
}

var (
	_ handlers_interfaces.RoomChatHandler = &RoomChatHandler{}
	_ handlers_interfaces.SendChatHandler = &SendChatHandler{}
)

func newRoomChatHandler(usecase usecases_interfaces.ChatUsecase) *RoomChatHandler {
	return &RoomChatHandler{
		usecase: usecase,
	}
}

func newSendChatHandler(usecase usecases_interfaces.ChatUsecase) *SendChatHandler {
	return &SendChatHandler{
		usecase: usecase,
	}
}
