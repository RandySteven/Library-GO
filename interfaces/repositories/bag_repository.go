package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BagRepository interface {
	Repository[models.Bag]
	UnitOfWork
	FindBagByUser(ctx context.Context, userID uint64) (result []*models.Bag, err error)
}
