package usecases_interfaces

import "context"

type DevUsecase interface {
	CreateBucket(ctx context.Context, name string) error
	GetListBuckets(ctx context.Context) ([]string, error)
	MessageBrokerCheckerHealth(ctx context.Context) (string, error)
}
