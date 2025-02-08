package crons_client

import (
	"context"
	"log"
)

func (s *scheduler) RunAllJobs(ctx context.Context) error {
	log.Println("Running all jobs")
	s.cron.Start()
	for _, registerJob := range s.RegisteredJobs() {
		if err := registerJob(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *scheduler) StopAllJobs(ctx context.Context) error {
	log.Println("Stopping scheduler...")

	cronCtx := s.cron.Stop()
	select {
	case <-cronCtx.Done():
		log.Println("All cron jobs stopped gracefully")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
