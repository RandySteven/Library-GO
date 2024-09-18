package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type UserRepository interface {
	Repository[models.User]
	UnitOfWork
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
}
