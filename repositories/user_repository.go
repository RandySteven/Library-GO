package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type userRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (u *userRepository) Save(ctx context.Context, entity *models.User) (result *models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByID(ctx context.Context, id uint64) (result *models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(ctx context.Context, entity *models.User) (result *models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) SetTx(tx *sql.Tx) {
	u.tx = tx
}

func (u *userRepository) GetTx(ctx context.Context) *sql.Tx {
	return u.tx
}

var _ repositories_interfaces.UserRepository = &userRepository{}

func newUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
