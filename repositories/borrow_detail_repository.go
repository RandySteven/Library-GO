package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type borrowDetailRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *borrowDetailRepository) Save(ctx context.Context, entity *models.BorrowDetail) (result *models.BorrowDetail, err error) {
	result = entity
	id, err := utils.Save[models.BorrowDetail](ctx, b.InitTrigger(), queries.InsertBorrowDetailQuery, &entity.BorrowID, &entity.BookID)
	if err != nil {
		return nil, err
	}
	result.ID = *id
	return result, nil
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
	var trigger repositories_interfaces.Trigger = b.db
	if b.tx != nil {
		trigger = b.tx
	}
	return trigger
}

func (b *borrowDetailRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.tx = tx
	return nil
}

func (b *borrowDetailRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *borrowDetailRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *borrowDetailRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *borrowDetailRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

var _ repositories_interfaces.BorrowDetailRepository = &borrowDetailRepository{}

func newBorrowDetailRepository(db *sql.DB) *borrowDetailRepository {
	return &borrowDetailRepository{
		db: db,
	}
}
