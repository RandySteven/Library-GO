package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"mime/multipart"
)

type RoomUsecase interface {
	CreateRoom(ctx context.Context, request *requests.CreateRoomRequest, fileHeader *multipart.FileHeader) (result *responses.CreateRoomResponse, customErr *apperror.CustomError)
	GetAllRooms(ctx context.Context) (results []*responses.ListRoomResponse, customErr *apperror.CustomError)
	GetRoomByID(ctx context.Context, id uint64) (result *responses.RoomDetailResponse, customErr *apperror.CustomError)
	UploadRoomPhoto(ctx context.Context, request *requests.UploadRoomPhoto) (customErr *apperror.CustomError)
}
