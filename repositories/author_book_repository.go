package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
)

type authorBookRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (a *authorBookRepository) Save(ctx context.Context, entity *models.AuthorBook) (result *models.AuthorBook, err error) {
	id, err := utils.Save[models.AuthorBook](ctx, a.InitTrigger(), queries.InsertAuthorBookQuery, &entity.AuthorID, &entity.BookID)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return result, nil
}

func (a *authorBookRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = a.db
	if a.tx != nil {
		trigger = a.tx
	}
	return trigger
}

func (a *authorBookRepository) BeginTx(ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	a.tx = tx
	return nil
}

func (a *authorBookRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorBookRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorBookRepository) SetTx(tx *sql.Tx) {
	a.tx = tx
}

func (a *authorBookRepository) GetTx(ctx context.Context) *sql.Tx {
	return a.tx
}

func (a *authorBookRepository) FindAuthorBookByBookID(ctx context.Context, bookID uint64) (result []*models.AuthorBook, err error) {
	rows, err := a.InitTrigger().QueryContext(ctx, queries.SelectAuthorBookByBookIDQuery.ToString(), bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		authorBook := new(models.AuthorBook)
		err = rows.Scan(&authorBook.ID, &authorBook.AuthorID, &authorBook.BookID, &authorBook.CreatedAt, &authorBook.UpdatedAt, &authorBook.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, authorBook)
	}
	return result, nil
}

var _ repositories_interfaces.AuthorBookRepository = &authorBookRepository{}

func newAuthorBookRepository(db *sql.DB) *authorBookRepository {
	return &authorBookRepository{
		db: db,
	}
}
