package publishers_interface

import (
	"context"
)

type BookPublisher interface {
	PubBookToElastic(ctx context.Context, rID string, bookID uint64) (id string, err error)
}
