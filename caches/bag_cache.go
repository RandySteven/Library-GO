package caches

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/RandySteven/Library-GO/utils"
	"github.com/go-redis/redis/v8"
)

type bagCache struct {
	redis *redis.Client
}

func (b *bagCache) SetUserBagCache(ctx context.Context, userId uint64, books []*responses.BookBagResponse) error {
	return caches_client.SetMultiple[responses.BookBagResponse](ctx, b.redis, fmt.Sprintf(enums.UserBagKey, utils.HashID(userId)), books)
}

func (b *bagCache) GetUserBagCache(ctx context.Context, userId uint64) (result []*responses.BookBagResponse, err error) {
	return caches_client.GetMultiple[responses.BookBagResponse](ctx, b.redis, fmt.Sprintf(enums.UserBagKey, utils.HashID(userId)))
}

func (b *bagCache) Del(ctx context.Context, key string) (err error) {
	return caches_client.Del[responses.BookBagResponse](ctx, b.redis, fmt.Sprintf(enums.UserBagKey, utils.HashID(ctx.Value(enums.UserID).(uint64))))
}

var _ caches_interfaces.BagCache = &bagCache{}

func newBagCache(redis *redis.Client) *bagCache {
	return &bagCache{redis: redis}
}
