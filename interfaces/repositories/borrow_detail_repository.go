package repositories_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/entities/models"
)

type BorrowDetailRepository interface {
	Saver[models.BorrowDetail]
	UnitOfWork
	FindByBorrowID(ctx context.Context, borrowID uint64) (results []*models.BorrowDetail, err error)
	FindByBorrowIDAndBookID(ctx context.Context, borrowID uint64, bookID uint64) (result *models.BorrowDetail, err error)
	FindCurrReturnDate(ctx context.Context) (results []*models.BorrowDetail, err error)
	UpdateReturnDateByBorrowIDAndBookID(ctx context.Context, borrowID uint64, bookID uint64) (result *models.BorrowDetail, err error)
	FindByBookID(ctx context.Context, bookID uint64) (results []*models.BorrowDetail, err error)
}
