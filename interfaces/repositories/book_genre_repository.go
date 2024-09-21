package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BookGenreRepository interface {
	Repository[models.BookGenre]
	UnitOfWork
}
