package caches_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type BagCache interface {
	SetUserBagCache(ctx context.Context, userId uint64, books []*responses.BookBagResponse) error
	GetUserBagCache(ctx context.Context, userId uint64) (result []*responses.BookBagResponse, err error)
}
