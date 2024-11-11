package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type (
	RoomChatUserRepository interface {
		Saver[models.RoomChatUser]
		Finder[models.RoomChatUser]
		UnitOfWork
	}
)
