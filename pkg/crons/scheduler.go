package crons_client

import (
	"context"
	"github.com/RandySteven/Library-GO/schedulers"
	"github.com/robfig/cron/v3"
	"log"
	"os"
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
	log.Println("Running all jobs")
	s.cron.Start()
	if err := s.updateBorrowDetailStatus(ctx); err != nil {
		return err
	}
	if err := s.testSchedulerLog(ctx); err != nil {
		return err
	}
	return nil
}

func (s *scheduler) updateBorrowDetailStatus(ctx context.Context) error {
	return s.runScheduler(ctx, os.Getenv("SCHEDULER_UPDATE_BOOK_STATUS"), s.dependencies.schedulers.BorrowScheduler.UpdateBorrowDetailStatusToExpired)
}

func (s *scheduler) StopAllJobs(ctx context.Context) error {
	log.Println("Stopping scheduler...")

	// Gracefully stop cron jobs
	cronCtx := s.cron.Stop() // Returns a channel that closes once all running jobs are complete
	select {
	case <-cronCtx.Done():
		log.Println("All cron jobs stopped gracefully")
		return nil
	case <-ctx.Done():
		return ctx.Err() // Context timeout or cancellation
	}
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

func (s *scheduler) testSchedulerLog(ctx context.Context) error {
	return s.runScheduler(ctx, os.Getenv("SCHEDULER_LOG_TEST"), func(ctx context.Context) error {
		log.Println("scheduler log well")
		return nil
	})
}

var _ Job = &scheduler{}

func NewScheduler(schedulers *schedulers.Schedulers) *scheduler {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &scheduler{
		cron:         cron.New(cron.WithSeconds(), cron.WithLocation(jakartaTime)),
		dependencies: dependenciesUsecases{schedulers: schedulers},
	}
}
