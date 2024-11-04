package schedulers_interfaces

import "context"

type LoggerScheduler interface {
	UploadLoggerScheduler(ctx context.Context) error
}
