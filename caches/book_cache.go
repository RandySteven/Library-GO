package caches

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/go-redis/redis/v8"
)

type bookCache struct {
	redis *redis.Client
}

func (b *bookCache) Set(ctx context.Context, key string, value *models.Book) (err error) {
	return caches_client.Set[models.Book](ctx, b.redis, fmt.Sprintf(enums.BookKey, key), value)
}

func (b *bookCache) Get(ctx context.Context, key string) (value *models.Book, err error) {
	return caches_client.Get[models.Book](ctx, b.redis, fmt.Sprintf(enums.BookKey, key))
}

func (b *bookCache) SetMultiData(ctx context.Context, values []*models.Book) (err error) {
	return caches_client.SetMultiple[models.Book](ctx, b.redis, enums.BooksKey, values)
}

func (b *bookCache) GetMultiData(ctx context.Context) (values []*models.Book, err error) {
	return caches_client.GetMultiple[models.Book](ctx, b.redis, enums.BooksKey)
}

func (b *bookCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	err = b.Del(ctx, key)
	if err != nil {
		return nil, err
	}
	err = b.SetMultiData(ctx, update.([]*models.Book))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *bookCache) Del(ctx context.Context, key string) (err error) {
	return b.redis.Del(ctx, key).Err()
}

var _ caches_interfaces.BookCache = &bookCache{}

func newBookCache(redis *redis.Client) *bookCache {
	return &bookCache{
		redis: redis,
	}
}
