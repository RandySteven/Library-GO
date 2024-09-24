package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BookGenreRepository interface {
	Repository[models.BookGenre]
	UnitOfWork
	FindBookGenreByBookID(ctx context.Context, bookID uint64) (result []*models.BookGenre, err error)
}
