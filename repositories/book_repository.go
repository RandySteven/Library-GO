package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type bookRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (b *bookRepository) Save(ctx context.Context, entity *models.Book) (result *models.Book, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) FindByID(ctx context.Context, id uint64) (result *models.Book, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) FindAll(ctx context.Context, skip uint64, take uint64) ([]*models.Book, error) {
	//TODO implement me
	panic("implement me")
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
