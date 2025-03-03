package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type AuthorBookRepository interface {
	Saver[models.AuthorBook]
	//UnitOfWork
	FindAuthorBookByBookID(ctx context.Context, bookID uint64) (result []*models.AuthorBook, err error)
}
