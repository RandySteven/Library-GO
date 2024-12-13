package caches

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/redis/go-redis/v9"
)

type borrowCache struct {
	redis *redis.Client
}

func (b *borrowCache) Set(ctx context.Context, key string, value *responses.BorrowDetailResponse) (err error) {
	return caches_client.Set[responses.BorrowDetailResponse](ctx, b.redis, fmt.Sprintf(enums.BorrowKey, key), value)
}

func (b *borrowCache) Get(ctx context.Context, key string) (value *responses.BorrowDetailResponse, err error) {
	return caches_client.Get[responses.BorrowDetailResponse](ctx, b.redis, fmt.Sprintf(enums.BorrowKey, key))
}

func (b *borrowCache) SetMultiData(ctx context.Context, values []*responses.BorrowListResponse) (err error) {
	userId := ctx.Value(enums.UserID).(uint64)
	return caches_client.SetMultiple[responses.BorrowListResponse](ctx, b.redis, fmt.Sprintf(enums.BorrowsKey, userId), values)
}

func (b *borrowCache) GetMultiData(ctx context.Context) (values []*responses.BorrowListResponse, err error) {
	userId := ctx.Value(enums.UserID).(uint64)
	return caches_client.GetMultiple[responses.BorrowListResponse](ctx, b.redis, fmt.Sprintf(enums.BorrowsKey, userId))
}

func (b *borrowCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	err = caches_client.Del[any](ctx, b.redis, key)
	if err != nil {
		return nil, err
	}
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
