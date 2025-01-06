package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type RoleUserRepository interface {
	Saver[models.RoleUser]
	Finder[models.RoleUser]
	FindRoleUserByUserID(ctx context.Context, id uint64) (result *models.RoleUser, err error)
	UnitOfWork
}
