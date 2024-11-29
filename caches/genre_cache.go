package caches

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/redis/go-redis/v9"
)

type genreCache struct {
	redis *redis.Client
}

func (g *genreCache) SetMultiData(ctx context.Context, values []*responses.ListGenresResponse) (err error) {
	return caches_client.SetMultiple[responses.ListGenresResponse](ctx, g.redis, enums.GenresKey, values)
}

func (g *genreCache) GetMultiData(ctx context.Context) (values []*responses.ListGenresResponse, err error) {
	return caches_client.GetMultiple[responses.ListGenresResponse](ctx, g.redis, enums.GenresKey)
}

var _ caches_interfaces.GenreCache = &genreCache{}

func newGenreCache(redis *redis.Client) *genreCache {
	return &genreCache{
		redis: redis,
	}
}
