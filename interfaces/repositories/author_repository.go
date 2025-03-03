package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type AuthorRepository interface {
	Saver[models.Author]
	Finder[models.Author]
	//UnitOfWork
	FindSelectedAuthorsByID(ctx context.Context, id []uint64) (result []*models.Author, err error)
}
