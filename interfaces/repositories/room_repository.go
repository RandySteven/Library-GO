package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type RoomRepository interface {
	Saver[models.Room]
	Finder[models.Room]
	UnitOfWork
}
