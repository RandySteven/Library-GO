package caches

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	"github.com/go-redis/redis/v8"
)

type bagCache struct {
	redis *redis.Client
}

func (b *bagCache) SetUserBagCache(ctx context.Context, userId uint64, books []*responses.BookBagResponse) error {
	//TODO implement me
	panic("implement me")
}

func (b *bagCache) GetUserBagCache(ctx context.Context, userId uint64) (result []*responses.BookBagResponse, err error) {
	//TODO implement me
	panic("implement me")
}

var _ caches_interfaces.BagCache = &bagCache{}

func newBagCache(redis *redis.Client) *bagCache {
	return &bagCache{redis: redis}
}
