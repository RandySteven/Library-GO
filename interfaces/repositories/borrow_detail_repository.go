package repositories_interfaces

import "github.com/RandySteven/Library-GO/entities/models"

type BorrowDetailRepository interface {
	Repository[models.BorrowDetail]
}
