package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BorrowRepository interface {
	Repository[models.Borrow]
}
