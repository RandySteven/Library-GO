package crons_client

import (
	"context"
)

type Scheduler interface {
	RemoveTempFile(ctx context.Context) error
	UpdateBookStatus(ctx context.Context) error
}
