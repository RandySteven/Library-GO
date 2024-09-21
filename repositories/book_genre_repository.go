package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type bookGenreRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bookGenreRepository) Save(ctx context.Context, entity *models.BookGenre) (result *models.BookGenre, err error) {
	return
}

func (b *bookGenreRepository) FindByID(ctx context.Context, id uint64) (result *models.BookGenre, err error) {
	return
}

func (b *bookGenreRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.BookGenre, error) {
	//TODO implement me
	return nil, nil
}

func (b *bookGenreRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	return
}

func (b *bookGenreRepository) Update(ctx context.Context, entity *models.BookGenre) (result *models.BookGenre, err error) {
	//TODO implement me
	return
}

func (b *bookGenreRepository) InitTrigger() repositories_interfaces.Trigger {
	//TODO implement me
	return nil
}

func (b *bookGenreRepository) BeginTx(ctx context.Context) error {
	//TODO implement me
	return nil
}

func (b *bookGenreRepository) CommitTx(ctx context.Context) error {
	//TODO implement me
	return nil
}

func (b *bookGenreRepository) RollbackTx(ctx context.Context) error {
	//TODO implement me
	return nil
}

func (b *bookGenreRepository) SetTx(tx *sql.Tx) {
	//TODO implement me
	return
}

func (b *bookGenreRepository) GetTx(ctx context.Context) *sql.Tx {
	//TODO implement me
	return nil
}

var _ repositories_interfaces.BookGenreRepository = &bookGenreRepository{}

func newBookGenreRepository(db *sql.DB) *bookGenreRepository {
	return &bookGenreRepository{
		db: db,
	}
}
