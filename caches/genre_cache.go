package caches

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	"github.com/go-redis/redis/v8"
)

type genreCache struct {
	redis *redis.Client
}

func (g *genreCache) SetMultiData(ctx context.Context, values []*responses.ListGenresResponse) (err error) {
	//TODO implement me
	panic("implement me")
}

func (g *genreCache) GetMultiData(ctx context.Context) (values []*responses.ListGenresResponse, err error) {
	//TODO implement me
	panic("implement me")
}

var _ caches_interfaces.GenreCache = &genreCache{}

func newGenreCache(redis *redis.Client) *genreCache {
	return &genreCache{
		redis: redis,
	}
}
