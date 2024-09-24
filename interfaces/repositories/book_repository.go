package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BookRepository interface {
	Repository[models.Book]
	UnitOfWork
	FindSelectedBooksId(ctx context.Context, bookIds []uint64) (result []*models.Book, err error)
}
