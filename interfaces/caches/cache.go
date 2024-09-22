package caches_interfaces

import (
	"context"
)

type Cache[T any] interface {
	Set(ctx context.Context, key string, value *T) (err error)
	Get(ctx context.Context, key string) (value *T, err error)
	SetMultiData(ctx context.Context, values []*T) (err error)
	GetMultiData(ctx context.Context) (values []*T, err error)
	Refresh(ctx context.Context, key string, update any) (value any, err error)
	Del(ctx context.Context, key string) (err error)
}
