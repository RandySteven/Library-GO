package caches_interfaces

import (
	"context"
)

type Cache[S any, M any] interface {
	SingleDataCache[S]
	MultiDataCache[M]
	Refresh(ctx context.Context, key string, update any) (value any, err error)
	Del(ctx context.Context, key string) (err error)
}

type SingleDataCache[T any] interface {
	Set(ctx context.Context, key string, value *T) (err error)
	Get(ctx context.Context, key string) (value *T, err error)
}

type MultiDataCache[T any] interface {
	SetMultiData(ctx context.Context, values []*T) (err error)
	GetMultiData(ctx context.Context) (values []*T, err error)
}
