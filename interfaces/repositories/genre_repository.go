package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type GenreRepository interface {
	Repository[models.Genre]
	UnitOfWork
	FindSelectedGenresByID(ctx context.Context, id []uint64) (result []*models.Genre, err error)
}
