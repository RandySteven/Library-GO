package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type GenreRepository interface {
	Repository[models.Genre]
	UnitOfWork
}
