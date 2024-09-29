package crons_client

import (
	"context"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"github.com/robfig/cron/v3"
)

type (
	scheduler struct {
		cron         *cron.Cron
		dependencies dependenciesUsecases
	}

	dependenciesUsecases struct {
		borrow usecases_interfaces.BorrowUsecase
	}
)

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
