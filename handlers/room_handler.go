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

type RoomHandler struct {
	usecase usecases_interfaces.RoomUsecase
}

func (r2 *RoomHandler) AddNewRoom(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CreateRoomRequest{}
		dataKey = `result`
	)
	if err := utils.BindRequest(r, &request); err != nil {
		return
	}
	thumbnailFile, fileHeader, err := r.FormFile("thumbnail")
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `failed to upload image`, nil, nil, err)
		return
	}
	defer thumbnailFile.Close()
	request.Thumbnail = thumbnailFile
	result, customErr := r2.usecase.CreateRoom(ctx, request, fileHeader)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to upload image`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to add new room`, &dataKey, result, nil)
}

func (r2 *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `rooms`
	)
	result, customErr := r2.usecase.GetAllRooms(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get data`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to get rooms`, &dataKey, result, nil)
}

func (r2 *RoomHandler) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 *RoomHandler) UploadRoom(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers_interfaces.RoomHandler = &RoomHandler{}

func newRoomHandler(usecase usecases_interfaces.RoomUsecase) *RoomHandler {
	return &RoomHandler{
		usecase: usecase,
	}
}
