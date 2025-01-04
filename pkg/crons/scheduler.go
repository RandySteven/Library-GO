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

func deleteCreateDir(dirName string) error {
	err := os.RemoveAll(dirName)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduler) refereshBookList(ctx context.Context) error {
	return s.runScheduler(ctx, os.Getenv("SCHEDULER_REFRESH_BOOK_CACHE"), func(ctx context.Context) error {
		return s.dependencies.schedulers.BookScheduler.RefreshBooksCache(ctx)
	})
}

func (s *scheduler) deleteStoryFiles(ctx context.Context) error {
	return s.runScheduler(ctx, os.Getenv("SCHEDULER_DELETE_FILE"), func(ctx context.Context) error {
		return deleteCreateDir("./temp-stories")
	})
}

func (s *scheduler) deleteImageFiles(ctx context.Context) error {
	return s.runScheduler(ctx, os.Getenv("SCHEDULER_DELETE_FILE"), func(ctx context.Context) error {
		return deleteCreateDir("./temp-stories")
	})
}

func (s *scheduler) RunAllJobs(ctx context.Context) error {
	log.Println("Running all jobs")
	s.cron.Start()
	if err := s.updateBorrowDetailStatus(ctx); err != nil {
		return err
	}
	if err := s.testSchedulerLog(ctx); err != nil {
		return err
	}
	if err := s.refereshBookList(ctx); err != nil {
		return err
	}
	if err := s.uploadLogFile(ctx); err != nil {
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

func (s *scheduler) uploadLogFile(ctx context.Context) error {
	return s.runScheduler(ctx, os.Getenv("SCHEDULER_UPLOAD_LOG_FILE"), s.dependencies.schedulers.LoggerScheduler.UploadLoggerScheduler)
}

var _ Job = &scheduler{}

func NewScheduler(schedulers *schedulers.Schedulers) *scheduler {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &scheduler{
		cron:         cron.New(cron.WithSeconds(), cron.WithLocation(jakartaTime)),
		dependencies: dependenciesUsecases{schedulers: schedulers},
	}
}
