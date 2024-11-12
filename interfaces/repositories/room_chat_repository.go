package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type (
	RoomChatRepository interface {
		Saver[models.RoomChat]
		Finder[models.RoomChat]
		UnitOfWork
	}
)
