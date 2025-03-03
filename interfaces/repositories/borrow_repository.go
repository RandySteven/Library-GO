package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BorrowRepository interface {
	Saver[models.Borrow]
	Finder[models.Borrow]
	//UnitOfWork
	FindByReferenceID(ctx context.Context, referenceID string) (result *models.Borrow, err error)
	FindByUserId(ctx context.Context, userId uint64) (result []*models.Borrow, err error)
	FindByMultipleBorrowID(ctx context.Context, borrowIDs []uint64) (result []*models.Borrow, err error)
}
