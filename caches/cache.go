package caches

import "github.com/go-redis/redis/v8"

type Caches struct {
	redis *redis.Client
}

func NewCaches(redis *redis.Client) *Caches {
	return &Caches{}
}
