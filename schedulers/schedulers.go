package schedulers

import (
	"github.com/RandySteven/Library-GO/caches"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/repositories"
)

type Schedulers struct {
	BorrowScheduler schedulers_interfaces.BorrowScheduler
	BookScheduler   schedulers_interfaces.BookScheduler
	LoggerScheduler schedulers_interfaces.LoggerScheduler
}

func NewSchedulers(repo *repositories.Repositories, cache *caches.Caches, aws aws_client.AWS) *Schedulers {
	return &Schedulers{
		BorrowScheduler: newBorrowScheduler(repo.BorrowRepo, repo.BorrowDetailRepo, repo.BookRepo),
		BookScheduler:   newBookScheduler(repo.BookRepo, repo.RatingRepo, cache.BookCache),
		LoggerScheduler: newLoggingScheduler(aws),
	}
}
