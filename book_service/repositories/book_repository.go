package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/book_service/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/book_service/interfaces/repositories"
	"github.com/RandySteven/Library-GO/pkg/utils"
)

type bookRepository struct {
	db *sql.DB
}

func (b *bookRepository) Create(ctx context.Context, request *models.Book) (uint64, error) {
	return utils.Save[models.Book]()
}

func (b *bookRepository) FindByID(ctx context.Context, id uint64) (*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) FindAll(ctx context.Context) ([]*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) Update(ctx context.Context, models *models.Book) (*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookRepository) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

var _ repositories_interfaces.BookRepository = &bookRepository{}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}
