package schedulers

import (
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
	"github.com/RandySteven/Library-GO/repositories"
)

type Schedulers struct {
	BorrowScheduler schedulers_interfaces.BorrowScheduler
}

func NewSchedulers(repo *repositories.Repositories) *Schedulers {
	return &Schedulers{
		BorrowScheduler: newBorrowScheduler(repo.BorrowRepo, repo.BorrowDetailRepo, repo.BookRepo),
	}
}
