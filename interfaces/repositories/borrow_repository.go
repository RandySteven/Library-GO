package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BorrowRepository interface {
	Repository[models.Borrow]
	UnitOfWork
	FindByReferenceID(ctx context.Context, referenceID string) (result *models.Borrow, err error)
}
