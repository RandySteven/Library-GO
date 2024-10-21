package schedulers_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BookScheduler interface {
	RefreshBooksCache(ctx context.Context) (books *models.Book, err error)
}
