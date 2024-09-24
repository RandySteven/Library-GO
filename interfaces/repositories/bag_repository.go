package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BagRepository interface {
	Repository[models.Bag]
	UnitOfWork
}
