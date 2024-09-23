package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type bookRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bookRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = b.db
	if b.tx != nil {
		trigger = b.tx
	}
	return trigger
}

func (b *bookRepository) BeginTx(ctx context.Context) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	b.tx = tx
	return nil
}

func (b *bookRepository) CommitTx(ctx context.Context) error {
	return b.tx.Commit()
}

func (b *bookRepository) RollbackTx(ctx context.Context) error {
	return b.tx.Rollback()
}

func (b *bookRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *bookRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *bookRepository) Save(ctx context.Context, entity *models.Book) (result *models.Book, err error) {
	id, err := utils.Save[models.Book](ctx, b.InitTrigger(), queries.InsertBookQuery, entity.Title, entity.Description, entity.Image, entity.Status)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bookRepository) FindByID(ctx context.Context, id uint64) (result *models.Book, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Book, error) {
	return utils.FindAll[models.Book](ctx, b.InitTrigger(), queries.SelectBooksQuery)
}

func (b *bookRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) Update(ctx context.Context, entity *models.Book) (result *models.Book, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.BookRepository = &bookRepository{}

func newBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}
