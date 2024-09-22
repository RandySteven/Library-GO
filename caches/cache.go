package caches

import "github.com/go-redis/redis/v8"

type Caches struct {
}

func NewCaches(redis *redis.Client) *Caches {
	return &Caches{}
}
