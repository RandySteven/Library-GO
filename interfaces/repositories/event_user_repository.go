package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type EventUserRepository interface {
	Saver[models.EventUser]
	Finder[models.EventUser]
	UnitOfWork
}
