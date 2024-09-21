package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type authorBookRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (a *authorBookRepository) Save(ctx context.Context, entity *models.AuthorBook) (result *models.AuthorBook, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorBookRepository) FindByID(ctx context.Context, id uint64) (result *models.AuthorBook, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorBookRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.AuthorBook, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorBookRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorBookRepository) Update(ctx context.Context, entity *models.AuthorBook) (result *models.AuthorBook, err error) {
	//TODO implement me
	panic("implement me")
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

var _ repositories_interfaces.AuthorBookRepository = &authorBookRepository{}

func newAuthorBookRepository(db *sql.DB) *authorBookRepository {
	return &authorBookRepository{
		db: db,
	}
}
