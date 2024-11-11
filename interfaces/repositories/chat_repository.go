package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type (
	ChatRepository interface {
		Saver[models.Chat]
		Finder[models.Chat]
		UnitOfWork
	}
)
