package caches

import (
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	"github.com/go-redis/redis/v8"
)

type Caches struct {
	BookCache   caches_interfaces.BookCache
	BorrowCache caches_interfaces.BorrowCache
	GenreCache  caches_interfaces.GenreCache
	redis       *redis.Client
}

func NewCaches(redis *redis.Client) *Caches {
	return &Caches{
		BookCache:   newBookCache(redis),
		BorrowCache: newBorrowCache(redis),
		GenreCache:  newGenreCache(redis),
		redis:       redis,
	}
}
