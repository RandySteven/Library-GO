package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
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
	id, err := utils.Save[models.Borrow](ctx, b.InitTrigger(), queries.InsertBorrowQuery, &entity.UserID, &entity.BorrowReference)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

func (b *borrowRepository) FindByID(ctx context.Context, id uint64) (result *models.Borrow, err error) {
	result = &models.Borrow{}
	err = utils.FindByID[models.Borrow](ctx, b.InitTrigger(), queries.SelectBorrowByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
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

func (b *borrowRepository) FindByReferenceID(ctx context.Context, referenceID string) (result *models.Borrow, err error) {
	result = &models.Borrow{}
	err = b.InitTrigger().QueryRowContext(ctx, queries.SelectBorrowByReference.ToString(), referenceID).
		Scan(
			&result.ID,
			&result.UserID,
			&result.BorrowReference,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *borrowRepository) FindByUserId(ctx context.Context, userId uint64) (result []*models.Borrow, err error) {
	rows, err := b.InitTrigger().QueryContext(ctx, queries.SelectBorrowUserIdQuery.ToString(), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		borrow := &models.Borrow{}
		err = rows.Scan(
			&borrow.ID,
			&borrow.UserID,
			&borrow.BorrowReference,
			&borrow.CreatedAt,
			&borrow.UpdatedAt,
			&borrow.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, borrow)
	}

	return result, nil
}

var _ repositories_interfaces.BorrowRepository = &borrowRepository{}

func newBorrowRepository(db *sql.DB) *borrowRepository {
	return &borrowRepository{
		db: db,
	}
}
