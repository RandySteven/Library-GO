package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type (
	ChatUsecase interface {
		CreateRoomChat(ctx context.Context, request *requests.CreateChatRoom) (result *responses.CreateRoomChatResponse, customErr *apperror.CustomError)
		ListRooms(ctx context.Context) (result []*responses.ListRoomChatsResponse, customErr *apperror.CustomError)
		GetRoomDetail(ctx context.Context, roomId uint64) (result *responses.RoomChatsResponse, customErr *apperror.CustomError)
		SendChat(ctx context.Context, request *requests.SendChat) (result *responses.ChatResponse, customErr *apperror.CustomError)
	}
)
