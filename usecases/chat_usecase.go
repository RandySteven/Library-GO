package usecases

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/RandySteven/Library-GO/utils"
	"log"
	"time"
)

type chatUsecase struct {
	chatRepo         repositories_interfaces.ChatRepository
	roomChatRepo     repositories_interfaces.RoomChatRepository
	roomChatUserRepo repositories_interfaces.RoomChatUserRepository
}

func (c *chatUsecase) refreshTx(tx *sql.Tx) {
	c.chatRepo.SetTx(tx)
	c.roomChatRepo.SetTx(tx)
	c.roomChatUserRepo.SetTx(tx)
}

func (c *chatUsecase) CreateRoomChat(ctx context.Context, request *requests.CreateChatRoom) (result *responses.CreateRoomChatResponse, customErr *apperror.CustomError) {
	//begin tx
	err := c.roomChatRepo.BeginTx(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to init transaction", err)
	}
	defer func() {
		if r := recover(); r != nil {
			_ = c.roomChatRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = c.roomChatRepo.RollbackTx(ctx)
		} else if err = c.roomChatRepo.CommitTx(ctx); err != nil {
			log.Println("failed to commit transaction:", err)
		}
		c.roomChatRepo.SetTx(nil)
	}()
	c.refreshTx(c.roomChatRepo.GetTx(ctx))

	roomChat, err := c.roomChatRepo.Save(ctx, &models.RoomChat{
		RoomName: request.RoomName,
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create room chat`, err)
	}

	for _, userId := range request.InviteUserIDs {
		_, err := c.roomChatUserRepo.Save(ctx, &models.RoomChatUser{
			RoomChatID: roomChat.ID,
			UserID:     userId,
		})
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create room chat user`, err)
		}
	}

	result = &responses.CreateRoomChatResponse{
		ID:        utils.HashID(roomChat.ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return result, nil
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
