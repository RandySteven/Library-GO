package caches_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type BagCache interface {
	SetUserBagCache(ctx context.Context, userId uint64, books []*responses.BookBagResponse) error
	GetUserBagCache(ctx context.Context, userId uint64) (result []*responses.BookBagResponse, err error)
	SetBookBagCache(ctx context.Context, bookBagCache *models.BookBagCache) (err error)
	GetBookBagCache(ctx context.Context, bookId string) (result *models.BookBagCache, err error)
	Del(ctx context.Context, key string) (err error)
}
