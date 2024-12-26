package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type UserRepository interface {
	Saver[models.User]
	Finder[models.User]
	UnitOfWork
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
	FindSelectedUsersByID(ctx context.Context, ids []uint64) (result []*models.User, err error)
}
