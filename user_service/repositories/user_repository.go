package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/user_service/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/user_service/interfaces/repositories"
)

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(ctx context.Context, models *models.User) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(ctx context.Context, models *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *sql.DB) repositories_interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

var _ repositories_interfaces.UserRepository = &userRepository{}
