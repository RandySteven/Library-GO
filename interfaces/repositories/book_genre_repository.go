package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BookGenreRepository interface {
	Saver[models.BookGenre]
	//UnitOfWork
	FindBookGenreByBookID(ctx context.Context, bookID uint64) (result []*models.BookGenre, err error)
	FindBookGenreByGenreID(ctx context.Context, genreID uint64) (result []*models.BookGenre, err error)
}
