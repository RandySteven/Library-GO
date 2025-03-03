package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type EventRepository interface {
	Saver[models.Event]
	Finder[models.Event]
	Updater[models.Event]
	//UnitOfWork
}
