package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type returnUsecase struct {
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository
	bookRepo         repositories_interfaces.BookRepository
	userRepo         repositories_interfaces.UserRepository
	transaction      repositories_interfaces.Transaction
}

func (r *returnUsecase) ReturnBook(ctx context.Context, request *requests.ReturnRequest) (result *responses.ReturnBooksResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)

	customErr = r.transaction.RunInTx(ctx, func(ctx context.Context) *apperror.CustomError {
		borrow, err := r.borrowRepo.FindByReferenceID(ctx, request.BorrowID)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow record`, err)
		}
		result = &responses.ReturnBooksResponse{
			BorrowID: borrow.ID,
			UserID:   userId,
		}
		for _, bookId := range request.BookIDs {
			_, err := r.borrowDetailRepo.UpdateReturnDateByBorrowIDAndBookID(ctx, borrow.ID, bookId)
			if err != nil {
				return apperror.NewCustomError(apperror.ErrInternalServer, `failed to get update detail`, err)
			}
			borrowDetail, err := r.borrowDetailRepo.FindByBorrowIDAndBookID(ctx, borrow.ID, bookId)
			if err != nil {
				return apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow detail`, err)
			}
			err = r.bookRepo.UpdateBookStatus(ctx, bookId, enums.Available)
			if err != nil {
				return apperror.NewCustomError(apperror.ErrInternalServer, `failed to update book status`, err)
			}
			result.ReturnedBooks = append(result.ReturnedBooks, &responses.ReturnBookResponse{
				BookID:             borrowDetail.BookID,
				VerifiedReturnDate: borrowDetail.VerifiedReturnDate,
			})
		}
		return nil
	})
	if customErr != nil {
		return nil, customErr
	}

	return result, nil
}

var _ usecases_interfaces.ReturnUsecase = &returnUsecase{}

func newReturnUsecase(
	borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	bookRepo repositories_interfaces.BookRepository,
	userRepo repositories_interfaces.UserRepository,
	transaction repositories_interfaces.Transaction) *returnUsecase {
	return &returnUsecase{
		borrowRepo:       borrowRepo,
		borrowDetailRepo: borrowDetailRepo,
		bookRepo:         bookRepo,
		userRepo:         userRepo,
		transaction:      transaction,
	}
}
