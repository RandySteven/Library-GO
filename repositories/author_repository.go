package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/utils"
)

type authorRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (a *authorRepository) InitTrigger() repositories_interfaces.Trigger {
	var trigger repositories_interfaces.Trigger = a.db
	if a.tx != nil {
		trigger = a.tx
	}
	return trigger
}

func (a *authorRepository) Save(ctx context.Context, entity *models.Author) (result *models.Author, err error) {
	id, err := utils.Save[models.Author](ctx, a.InitTrigger(), ``, &entity.Name, &entity.Nationality)
	if err != nil {
		return nil, err
	}
	result = entity
	result.ID = *id
	return
}

func (a *authorRepository) FindByID(ctx context.Context, id uint64) (result *models.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) Update(ctx context.Context, entity *models.Author) (result *models.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) BeginTx(ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	a.tx = tx
	return nil
}

func (a *authorRepository) CommitTx(ctx context.Context) error {
	return a.tx.Commit()
}

func (a *authorRepository) RollbackTx(ctx context.Context) error {
	return a.tx.Rollback()
}

func (a *authorRepository) SetTx(tx *sql.Tx) {
	a.tx = tx
}

func (a *authorRepository) GetTx(ctx context.Context) *sql.Tx {
	return a.tx
}

var _ repositories_interfaces.AuthorRepository = &authorRepository{}

func newAuthorRepository(db *sql.DB) *authorRepository {
	return &authorRepository{
		db: db,
	}
}
