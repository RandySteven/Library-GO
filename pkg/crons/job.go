package crons_client

import (
	"context"
)

type Job interface {
	RunAllJobs(ctx context.Context) error
	testSchedulerLog(ctx context.Context) error
	updateBorrowDetailStatus(ctx context.Context) error
	StopAllJobs(ctx context.Context) error
}
