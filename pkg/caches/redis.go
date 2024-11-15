package caches_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var redisTimeout = os.Getenv("REDIS_EXPIRATION")

type RedisClient struct {
	client  *redis.Client
	limiter *redis_rate.Limiter
}

func NewRedisCache(config *configs.Config) (*RedisClient, error) {
	redisCfg := config.Config.Redis
	addr := fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port)
	log.Println("connecting to redis : ", addr)
	opt, _ := redis.ParseURL(fmt.Sprintf(`rediss://default:%s@%s:%s`, redisCfg.Password, redisCfg.Host, redisCfg.Port))

	client := redis.NewClient(opt)
	limiter := redis_rate.NewLimiter(client)
	return &RedisClient{
		client:  client,
		limiter: limiter,
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

func (c *RedisClient) RateLimiter(ctx context.Context) error {
	clientIP := ctx.Value(enums.ClientIP).(string)
	res, err := c.limiter.Allow(ctx, clientIP, redis_rate.PerMinute(10))
	if err != nil {
		return err
	}
	if res.Remaining == 0 {
		return errors.New("Rate limit exceeded")
	}
	return nil
}
