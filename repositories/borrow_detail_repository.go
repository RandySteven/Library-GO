package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type borrowDetailRepository struct {
	db *sql.DB
}

func (b *borrowDetailRepository) Save(ctx context.Context, entity *models.BorrowDetail) (result *models.BorrowDetail, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) FindByID(ctx context.Context, id uint64) (result *models.BorrowDetail, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.BorrowDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) Update(ctx context.Context, entity *models.BorrowDetail) (result *models.BorrowDetail, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) InitTrigger() repositories_interfaces.Trigger {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) SetTx(tx *sql.Tx) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowDetailRepository) GetTx(ctx context.Context) *sql.Tx {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.BorrowDetailRepository = &borrowDetailRepository{}

func newBorrowDetailRepository(db *sql.DB) *borrowDetailRepository {
	return &borrowDetailRepository{
		db: db,
	}
}
