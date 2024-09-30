package crons_client

import (
	"context"
	"github.com/RandySteven/Library-GO/schedulers"
	"github.com/robfig/cron/v3"
	"time"
)

type (
	scheduler struct {
		cron         *cron.Cron
		dependencies dependenciesUsecases
	}

	dependenciesUsecases struct {
		schedulers *schedulers.Schedulers
	}
)

func (s *scheduler) RunAllJobs(ctx context.Context) error {
	if err := s.updateBorrowDetailStatus(ctx); err != nil {
		return err
	}
	return nil
}

func (s *scheduler) updateBorrowDetailStatus(ctx context.Context) error {
	return s.runScheduler(ctx, "@daily", s.dependencies.schedulers.BorrowScheduler.UpdateBorrowDetailStatusToExpired)
}

func (s *scheduler) runScheduler(ctx context.Context, spec string, schedulerFunc func(ctx context.Context) error) error {
	_, err := s.cron.AddFunc(spec, func() {
		err := schedulerFunc(ctx)
		if err != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	return nil
}

var _ Job = &scheduler{}

func NewScheduler(schedulers *schedulers.Schedulers) *scheduler {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &scheduler{
		cron:         cron.New(cron.WithLocation(jakartaTime)),
		dependencies: dependenciesUsecases{schedulers: schedulers},
	}
}
