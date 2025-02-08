package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type bookRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bookRepository) Trigger() repositories_interfaces.Trigger {
	return utils.InitTrigger(b.db, b.tx)
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
	return utils.CommitTx(ctx, b.tx)
}

func (b *bookRepository) RollbackTx(ctx context.Context) error {
	return utils.RollbackTx(ctx, b.tx)
}

func (b *bookRepository) SetTx(tx *sql.Tx) {
	b.tx = tx
}

func (b *bookRepository) GetTx(ctx context.Context) *sql.Tx {
	return b.tx
}

func (b *bookRepository) Save(ctx context.Context, entity *models.Book) (result *models.Book, err error) {
	id, err := utils.Save[models.Book](ctx, b.Trigger(), queries.InsertBookQuery, entity.Title, entity.Description, entity.Image, entity.Status)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (b *bookRepository) FindByID(ctx context.Context, id uint64) (result *models.Book, err error) {
	result = &models.Book{}
	err = utils.FindByID[models.Book](ctx, b.Trigger(), queries.SelectBookByIDQuery, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *bookRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Book, error) {
	if skip == 0 && take == 0 {
		return utils.FindAll[models.Book](ctx, b.Trigger(), queries.SelectBooksQuery)
	}
	if skip == 1 {
		skip = 0
	} else {
		skip = skip*take - take
	}
	books := []*models.Book{}
	rows, err := b.Trigger().QueryContext(ctx, queries.SelectBookPaginateQuery.ToString(), take, skip)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.Image,
			&book.Status, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b *bookRepository) FindSelectedBooksId(ctx context.Context, ids []uint64) (results []*models.Book, err error) {
	selectStr := utils.SelectIdIn(queries.SelectBooksQuery, ids)
	rows, err := b.db.QueryContext(ctx, selectStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := new(models.Book)
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image, &book.Status, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, book)
	}
	return results, nil
}

func (b *bookRepository) FindBookStatus(ctx context.Context, id uint64, status enums.BookStatus) (isExist bool, err error) {
	result := &models.Book{}
	err = b.Trigger().QueryRowContext(ctx, queries.SelectBookAndStatus.ToString(), id, status).Scan(
		&result.ID, &result.Title, &result.Description, &result.Image, &result.Status, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return false, err
	}
	return result.Status == status, nil
}

func (b *bookRepository) UpdateBookStatus(ctx context.Context, id uint64, status enums.BookStatus) error {
	_, err := b.Trigger().ExecContext(ctx, queries.UpdateBookStatusQuery.ToString(), status, id)
	if err != nil {
		return err
	}
	return nil
}

var _ repositories_interfaces.BookRepository = &bookRepository{}

func newBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}
