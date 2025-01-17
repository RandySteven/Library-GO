package caches

import (
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	"github.com/redis/go-redis/v9"
)

type Caches struct {
	BookCache   caches_interfaces.BookCache
	BorrowCache caches_interfaces.BorrowCache
	GenreCache  caches_interfaces.GenreCache
	BagCache    caches_interfaces.BagCache
	redis       *redis.Client
}

func NewCaches(redis *redis.Client) *Caches {
	return &Caches{
		BookCache:   newBookCache(redis),
		BorrowCache: newBorrowCache(redis),
		GenreCache:  newGenreCache(redis),
		BagCache:    newBagCache(redis),
		redis:       redis,
	}
}
