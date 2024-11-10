package usecases_interfaces

import "context"

type EventUsecase interface {
	CreateEvent(ctx context.Context)
}
