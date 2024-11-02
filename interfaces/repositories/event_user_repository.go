package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type EventUserRepository interface {
	Repository[models.EventUser]
	UnitOfWork
}
