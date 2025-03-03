package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type (
	ChatRepository interface {
		Saver[models.Chat]
		Finder[models.Chat]
		//UnitOfWork
		FindChatByRoomID(ctx context.Context, roomChatID uint64) (result []*models.Chat, err error)
	}
)
