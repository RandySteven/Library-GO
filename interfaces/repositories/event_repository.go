package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type EventRepository interface {
	Repository[models.Event]
	UnitOfWork
}
