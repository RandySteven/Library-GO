package schedulers

import (
	"github.com/RandySteven/Library-GO/caches"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
	"github.com/RandySteven/Library-GO/repositories"
)

type Schedulers struct {
	BorrowScheduler schedulers_interfaces.BorrowScheduler
	BookScheduler   schedulers_interfaces.BookScheduler
}

func NewSchedulers(repo *repositories.Repositories, cache *caches.Caches) *Schedulers {
	return &Schedulers{
		BorrowScheduler: newBorrowScheduler(repo.BorrowRepo, repo.BorrowDetailRepo, repo.BookRepo),
		BookScheduler:   newBookScheduler(repo.BookRepo, repo.RatingRepo, cache.BookCache),
	}
}
