package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
)

type BookRepository interface {
	Saver[models.Book]
	Finder[models.Book]
	UnitOfWork
	FindSelectedBooksId(ctx context.Context, bookIds []uint64) (result []*models.Book, err error)
	FindBookStatus(ctx context.Context, id uint64, status enums.BookStatus) (isExist bool, err error)
	UpdateBookStatus(ctx context.Context, id uint64, status enums.BookStatus) error
}
