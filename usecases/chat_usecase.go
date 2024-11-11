package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type chatUsecase struct {
	chatRepo         repositories_interfaces.ChatRepository
	roomChatRepo     repositories_interfaces.RoomChatRepository
	roomChatUserRepo repositories_interfaces.RoomChatUserRepository
}

func (c *chatUsecase) CreateRoomChat(ctx context.Context, request *requests.CreateChatRoom) (result *responses.CreateRoomChatResponse, customErr *apperror.CustomError) {

	return
}

func (c *chatUsecase) ListRooms(ctx context.Context) (result []*responses.ListRoomChatsResponse, customErr *apperror.CustomError) {
	return
}

func (c *chatUsecase) GetRoomDetail(ctx context.Context, roomId uint64) (result *responses.RoomChatsResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.ChatUsecase = &chatUsecase{}

func newChatUsecase(
	chatRepo repositories_interfaces.ChatRepository,
	roomChatRepo repositories_interfaces.RoomChatRepository,
	roomChatUserRepo repositories_interfaces.RoomChatUserRepository) *chatUsecase {
	return &chatUsecase{
		chatRepo:         chatRepo,
		roomChatRepo:     roomChatRepo,
		roomChatUserRepo: roomChatUserRepo,
	}
}
