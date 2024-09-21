package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type borrowRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *borrowRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = b.db
	if b.tx != nil {
		trigger = b.tx
	}
	return trigger
}

func (b *borrowRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.SetTx(tx)
	return nil
}

func (b *borrowRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *borrowRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *borrowRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *borrowRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *borrowRepository) Save(ctx context.Context, entity *models.Borrow) (result *models.Borrow, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowRepository) FindByID(ctx context.Context, id uint64) (result *models.Borrow, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Borrow, error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowRepository) Update(ctx context.Context, entity *models.Borrow) (result *models.Borrow, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.BorrowRepository = &borrowRepository{}

func newBorrowRepository(db *sql.DB) *borrowRepository {
	return &borrowRepository{
		db: db,
	}
}
