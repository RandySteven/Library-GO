package caches

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/redis/go-redis/v9"
	"time"
)

type bookCache struct {
	redis *redis.Client
}

func (b *bookCache) Set(ctx context.Context, key string, value *responses.BookDetailResponse) (err error) {
	return caches_client.Set[responses.BookDetailResponse](ctx, b.redis, fmt.Sprintf(enums.BookKey, key), value)
}

func (b *bookCache) Get(ctx context.Context, key string) (value *responses.BookDetailResponse, err error) {
	return caches_client.Get[responses.BookDetailResponse](ctx, b.redis, fmt.Sprintf(enums.BookKey, key))
}

func (b *bookCache) SetMultiData(ctx context.Context, values []*responses.ListBooksResponse) (err error) {
	return caches_client.SetMultiple[responses.ListBooksResponse](ctx, b.redis, enums.BooksKey, values)
}

func (b *bookCache) GetMultiData(ctx context.Context) (values []*responses.ListBooksResponse, err error) {
	return caches_client.GetMultiple[responses.ListBooksResponse](ctx, b.redis, enums.BooksKey)
}

func (b *bookCache) Refresh(ctx context.Context, key string, update any) (value any, err error) {
	err = b.Del(ctx, key)
	if err != nil {
		return nil, err
	}
	//err = b.SetMultiData(ctx, update.([]*models.Book))
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (b *bookCache) Del(ctx context.Context, key string) (err error) {
	return b.redis.Del(ctx, key).Err()
}

func (b *bookCache) SetBookPage(ctx context.Context, page uint64, result []*responses.ListBooksResponse) error {
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}
	err = b.redis.Set(ctx, fmt.Sprintf(enums.BookPageKey, page), jsonData, time.Duration(enums.Expired)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (b *bookCache) GetBookPage(ctx context.Context, page uint64) (result []*responses.ListBooksResponse, err error) {
	val, err := b.redis.Get(ctx, fmt.Sprintf(enums.BookPageKey, page)).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(val, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ caches_interfaces.BookCache = &bookCache{}

func newBookCache(redis *redis.Client) *bookCache {
	return &bookCache{
		redis: redis,
	}
}
