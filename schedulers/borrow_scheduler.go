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
	//if err := b.borrowDetailRepo.BeginTx(ctx); err != nil {
	//	return err
	//}
	//borrowDetails, err := b.borrowDetailRepo.FindCurrReturnDate(ctx)
	//if err != nil {
	//	return err
	//}
	//var bookIds = []uint64{}
	//for _, borrowDetail := range borrowDetails {
	//	bookIds = append(bookIds, borrowDetail.BookID)
	//}
	//
	//for _, bookId := range bookIds {
	//	err = b.bookRepo.UpdateBookStatus(ctx, bookId, enums.Expired)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

var _ schedulers_interfaces.BorrowScheduler = &borrowScheduler{}

func newBorrowScheduler(borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	bookRepo repositories_interfaces.BookRepository) *borrowScheduler {
	return &borrowScheduler{borrowRepo, borrowDetailRepo, bookRepo}
}
