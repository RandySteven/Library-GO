package caches_client

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

var redisTimeout = os.Getenv("REDIS_EXPIRATION")

type RedisClient struct {
	client *redis.Client
}

func NewRedisCache(config *configs.Config) (*RedisClient, error) {
	redisCfg := config.Config.Redis
	addr := fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port)
	log.Println("connecting to redis : ", addr)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		MinIdleConns: redisCfg.MinIddleConns,
		PoolSize:     redisCfg.PoolSize,
		PoolTimeout:  time.Duration(redisCfg.PoolTimeout) * time.Second,
		DB:           0,
	})
	return &RedisClient{
		client: client,
	}, nil
}

func (c *RedisClient) Ping() error {
	return c.client.Ping(context.Background()).Err()
}

func (c *RedisClient) Client() *redis.Client {
	return c.client
}

func (c *RedisClient) ClearCache(ctx context.Context) error {
	return c.client.FlushDB(ctx).Err()
}
