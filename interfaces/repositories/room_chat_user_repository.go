package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type (
	RoomChatUserRepository interface {
		Saver[models.RoomChatUser]
		Finder[models.RoomChatUser]
		UnitOfWork
		FindUserRooms(ctx context.Context, userId uint64) (result []*models.RoomChatUser, err error)
	}
)
