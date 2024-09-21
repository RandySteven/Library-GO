package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BookRepository interface {
	Repository[models.Book]
	UnitOfWork
}
