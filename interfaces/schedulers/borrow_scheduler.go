package schedulers_interfaces

import "context"

type BorrowScheduler interface {
	UpdateBorrowDetailStatusToExpired(ctx context.Context) error
}
