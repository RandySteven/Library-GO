package schedulers

import (
	"context"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	schedulers_interfaces "github.com/RandySteven/Library-GO/interfaces/schedulers"
)

type borrowScheduler struct {
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository
	bookRepo         repositories_interfaces.BookRepository
}

func (b *borrowScheduler) UpdateBorrowDetailStatusToExpired(ctx context.Context) error {
	//preconditions:
	// - init transaction
	//1. get all borrow detail tables
	//2. get current date of returned_date
	//3. if returned_date == todays date
	//  a. return all book ids
	//4. update all book ids into expired
	return nil
}

var _ schedulers_interfaces.BorrowScheduler = &borrowScheduler{}

func newBorrowScheduler(borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	bookRepo repositories_interfaces.BookRepository) *borrowScheduler {
	return &borrowScheduler{borrowRepo, borrowDetailRepo, bookRepo}
}
