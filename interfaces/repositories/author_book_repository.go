package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type AuthorBookRepository interface {
	Repository[models.AuthorBook]
	UnitOfWork
}
