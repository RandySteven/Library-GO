package schedulers_interfaces

import (
	"context"
)

type BookScheduler interface {
	RefreshBooksCache(ctx context.Context) (err error)
}
