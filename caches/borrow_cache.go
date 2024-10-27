package caches

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/go-redis/redis/v8"
)

type borrowCache struct {
	redis *redis.Client
}

func (b *borrowCache) Set(ctx context.Context, key string, value *responses.BorrowDetailResponse) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowCache) Get(ctx context.Context, key string) (value *responses.BorrowDetailResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowCache) SetMultiData(ctx context.Context, values []*responses.BorrowListResponse) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowCache) GetMultiData(ctx context.Context) (values []*responses.BorrowListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *borrowCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	return
}

func (b *borrowCache) Del(ctx context.Context, key string) (err error) {
	return caches_client.Del[models.Borrow](ctx, b.redis, key)
}

var _ caches_interfaces.BorrowCache = &borrowCache{}

func newBorrowCache(redis *redis.Client) *borrowCache {
	return &borrowCache{
		redis: redis,
	}
}
