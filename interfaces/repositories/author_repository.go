package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type AuthorRepository interface {
	Repository[models.Author]
	UnitOfWork
}
