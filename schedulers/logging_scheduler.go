package schedulers

import (
	"context"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
)

type loggerScheduler struct {
}

func (l *loggerScheduler) UploadLoggerScheduler(ctx context.Context) error {
	return nil
}

var _ schedulers_interfaces.LoggerScheduler = &loggerScheduler{}

func newLoggingScheduler() *loggerScheduler {
	return &loggerScheduler{}
}
