package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type bagRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bagRepository) CheckBagExists(ctx context.Context, bag *models.Bag) (bool, error) {
	exists := 1
	err := b.InitTrigger().QueryRowContext(ctx, queries.SelectExistBookAlreadyInBag.ToString(), &bag.BookID, &bag.UserID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (b *bagRepository) FindBagByUser(ctx context.Context, userID uint64) (result []*models.Bag, err error) {
	rows, err := b.InitTrigger().QueryContext(ctx, queries.SelectBagByUserQuery.ToString(), userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		bag := new(models.Bag)
		err = rows.Scan(&bag.ID, &bag.UserID, &bag.BookID)
		if err != nil {
			return nil, err
		}
		result = append(result, bag)
	}
	return result, nil
}

func (b *bagRepository) Save(ctx context.Context, entity *models.Bag) (result *models.Bag, err error) {
	id, err := utils.Save[models.Bag](ctx, b.InitTrigger(), queries.InsertBagQuery, &entity.UserID, &entity.BookID)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bagRepository) FindByID(ctx context.Context, id uint64) (result *models.Bag, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bagRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Bag, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bagRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bagRepository) Update(ctx context.Context, entity *models.Bag) (result *models.Bag, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bagRepository) DeleteByUserAndBook(ctx context.Context, userId uint64, bookId uint64) error {
	_, err := b.InitTrigger().ExecContext(ctx, queries.DeleteByUserAndBookQuery.ToString(), userId, bookId)
	if err != nil {
		return err
	}
	return nil
}

func (b *bagRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = b.db
	if b.tx != nil {
		trigger = b.tx
	}
	return trigger
}

func (b *bagRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.tx = tx
	return nil
}

func (b *bagRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *bagRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *bagRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *bagRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *bagRepository) DeleteUserBag(ctx context.Context, userId uint64) error {
	_, err := b.InitTrigger().ExecContext(ctx, queries.DeleteUserBagQuery.ToString(), userId)
	if err != nil {
		return err
	}
	return nil
}

var _ repositories_interfaces.BagRepository = &bagRepository{}

func newBagRepository(db *sql.DB) *bagRepository {
	return &bagRepository{
		db: db,
	}
}
