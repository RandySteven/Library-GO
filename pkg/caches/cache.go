package caches_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func getRedisTimeout() time.Duration {
	redisDurrTime, _ := strconv.Atoi(redisTimeout)
	redisDurrTime64 := int64(redisDurrTime)
	return time.Duration(redisDurrTime64) * time.Second
}

func Set[T any](ctx context.Context, redis *redis.Client, key string, value *T) (err error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}
	return redis.Set(ctx, key, jsonData, getRedisTimeout()).Err()
}

func Get[T any](ctx context.Context, client *redis.Client, key string) (value *T, err error) {
	val, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(val, &value)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	return value, nil
}

func SetMultiple[T any](ctx context.Context, redis *redis.Client, key string, value []*T) (err error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}
	return redis.Set(ctx, key, jsonData, getRedisTimeout()).Err()
}

func GetMultiple[T any](ctx context.Context, redis *redis.Client, key string) (value []*T, err error) {
	val, err := redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(val, &value)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}
	return value, nil
}

func Del[T any](ctx context.Context, redis *redis.Client, key string) (err error) {
	return redis.Del(ctx, key).Err()
}
