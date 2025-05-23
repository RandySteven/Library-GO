package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/google/uuid"
	"sync"
)

type chatUsecase struct {
	chatRepo         repositories_interfaces.ChatRepository
	roomChatRepo     repositories_interfaces.RoomChatRepository
	roomChatUserRepo repositories_interfaces.RoomChatUserRepository
	userRepo         repositories_interfaces.UserRepository
}

func (c *chatUsecase) CreateRoomChat(ctx context.Context, request *requests.CreateChatRoom) (result *responses.CreateRoomChatResponse, customErr *apperror.CustomError) {
	//begin tx
	//err := c.roomChatRepo.BeginTx(ctx)
	//if err != nil {
	//	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to init transaction", err)
	//}
	//defer func() {
	//	if r := recover(); r != nil {
	//		_ = c.roomChatRepo.RollbackTx(ctx)
	//		panic(r)
	//	} else if customErr != nil {
	//		_ = c.roomChatRepo.RollbackTx(ctx)
	//	} else if err = c.roomChatRepo.CommitTx(ctx); err != nil {
	//		log.Println("failed to commit transaction:", err)
	//	}
	//	c.roomChatRepo.SetTx(nil)
	//}()
	//c.refreshTx(c.roomChatRepo.GetTx(ctx))
	//
	//userId := ctx.Value(enums.UserID).(uint64)
	//request.InviteUserIDs = append(request.InviteUserIDs, userId)
	//
	//users, err := c.userRepo.FindSelectedUsersByID(ctx, request.InviteUserIDs)
	//if err != nil {
	//	return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get list of users`, err)
	//}
	//
	//roomName := request.RoomName
	//if roomName != "" {
	//	for _, user := range users {
	//		roomName += user.Name + "|"
	//	}
	//}
	//
	//roomChat, err := c.roomChatRepo.Save(ctx, &models.RoomChat{
	//	RoomName: roomName,
	//})
	//
	//if err != nil {
	//	return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create room chat`, err)
	//}
	//
	//for _, user := range users {
	//	_, err := c.roomChatUserRepo.Save(ctx, &models.RoomChatUser{
	//		RoomChatID: roomChat.ID,
	//		UserID:     user.ID,
	//	})
	//	if err != nil {
	//		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create room chat user`, err)
	//	}
	//}
	//
	//result = &responses.CreateRoomChatResponse{
	//	ID:        utils.HashID(roomChat.ID),
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//}
	//
	//return result, nil
	return nil, nil
}

func (c *chatUsecase) ListRooms(ctx context.Context) (result []*responses.ListRoomChatsResponse, customErr *apperror.CustomError) {
	result = []*responses.ListRoomChatsResponse{}
	userId := ctx.Value(enums.UserID).(uint64)

	roomChatUsers, err := c.roomChatUserRepo.FindUserRooms(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `error find room chat users`, err)
	}

	for _, roomChatUser := range roomChatUsers {
		roomChat, err := c.roomChatRepo.FindByID(ctx, roomChatUser.RoomChatID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get room chat`, err)
		}
		result = append(result, &responses.ListRoomChatsResponse{
			RoomChatID:   roomChatUser.RoomChatID,
			RoomChatName: roomChat.RoomName,
		})
	}

	return result, nil
}

func (c *chatUsecase) findUsersInRoom(ctx context.Context, roomId uint64) {}

func (c *chatUsecase) GetRoomDetail(ctx context.Context, roomId uint64) (result *responses.RoomChatsResponse, customErr *apperror.CustomError) {
	return
}

func (c *chatUsecase) SendChat(ctx context.Context, request *requests.SendChat) (result *responses.ChatResponse, customErr *apperror.CustomError) {
	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
		_           = context.WithValue(ctx, enums.RequestID, uuid.NewString())
		mysqlCtx    = context.WithValue(ctx, enums.RequestID, uuid.NewString())
		chat        = &models.Chat{
			RoomChatID: request.RoomID,
			Chat:       request.Chat,
			UserID:     ctx.Value(enums.UserID).(uint64),
		}
	)

	wg.Add(2)

	//async redis
	go func() {
		defer wg.Done()

	}()

	//async db
	go func() {
		defer wg.Done()
		_, err := c.chatRepo.Save(mysqlCtx, chat)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to save chat at db`, err)
			return
		}
	}()

	go func() {
		wg.Wait()
	}()
	return
}

var _ usecases_interfaces.ChatUsecase = &chatUsecase{}

func newChatUsecase(
	chatRepo repositories_interfaces.ChatRepository,
	roomChatRepo repositories_interfaces.RoomChatRepository,
	roomChatUserRepo repositories_interfaces.RoomChatUserRepository,
	userRepo repositories_interfaces.UserRepository) *chatUsecase {
	return &chatUsecase{
		chatRepo:         chatRepo,
		roomChatRepo:     roomChatRepo,
		roomChatUserRepo: roomChatUserRepo,
		userRepo:         userRepo,
	}
}
