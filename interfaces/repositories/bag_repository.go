package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BagRepository interface {
	Saver[models.Bag]
	//UnitOfWork
	FindBagByUser(ctx context.Context, userID uint64) (result []*models.Bag, err error)
	CheckBagExists(ctx context.Context, bag *models.Bag) (bool, error)
	DeleteUserBag(ctx context.Context, userId uint64) error
	DeleteByUserAndSelectedBooks(ctx context.Context, userId uint64, bookId []uint64) error
}
