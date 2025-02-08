package crons_client

import (
	"context"
)

type (
	job func(ctx context.Context) error

	Job interface {
		RunAllJobs(ctx context.Context) error
		testSchedulerLog(ctx context.Context) error
		updateBorrowDetailStatus(ctx context.Context) error
		refereshBookList(ctx context.Context) error
		deleteStoryFiles(ctx context.Context) error
		deleteImageFiles(ctx context.Context) error
		uploadLogFile(ctx context.Context) error
		StopAllJobs(ctx context.Context) error
	}
)

func (s *scheduler) register(registerJob ...job) (jobs []job) {
	for _, job := range registerJob {
		jobs = append(jobs, job)
	}
	return jobs
}

func (s *scheduler) RegisteredJobs() (jobs []job) {
	jobs = s.register(
		s.testSchedulerLog,
		s.updateBorrowDetailStatus,
		s.refereshBookList,
		s.deleteStoryFiles,
		s.deleteImageFiles,
		s.uploadLogFile,
	)
	return jobs
}
