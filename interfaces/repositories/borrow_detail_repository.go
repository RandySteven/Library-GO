package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BorrowDetailRepository interface {
	Repository[models.BorrowDetail]
	UnitOfWork
	FindByBorrowID(ctx context.Context, borrowID uint64) (results []*models.BorrowDetail, err error)
}
