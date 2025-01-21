package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type RoomPhotoRepository interface {
	Saver[models.RoomPhoto]
	Finder[models.RoomPhoto]
	UnitOfWork
}
