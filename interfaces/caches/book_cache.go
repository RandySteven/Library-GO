package caches_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type BookCache interface {
	Cache[responses.BookDetailResponse, responses.ListBooksResponse]
	SetBookPage(ctx context.Context, page uint64, result []*responses.ListBooksResponse) error
	GetBookPage(ctx context.Context, page uint64) (result []*responses.ListBooksResponse, err error)
}
