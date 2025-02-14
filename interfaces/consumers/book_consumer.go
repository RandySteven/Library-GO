package consumers_interfaces

import (
	"context"
)

type BookConsumer interface {
	ConsumeBookToAddElastic(ctx context.Context) (err error)
	ConsumeBookAddToRedis(ctx context.Context) (err error)
}
