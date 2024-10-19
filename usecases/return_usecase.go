package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"log"
)

type returnUsecase struct {
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository
	bookRepo         repositories_interfaces.BookRepository
	userRepo         repositories_interfaces.UserRepository
}

func (r *returnUsecase) refreshTx(ctx context.Context) {
	r.borrowDetailRepo.SetTx(nil)
	r.setTx(ctx)
}

func (r *returnUsecase) setTx(ctx context.Context) {
	tx := r.borrowDetailRepo.GetTx(ctx)
	r.borrowRepo.SetTx(tx)
	r.bookRepo.SetTx(tx)
	r.userRepo.SetTx(tx)
}

func (r *returnUsecase) ReturnBook(ctx context.Context, request *requests.ReturnRequest) (result *responses.ReturnBooksResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	err := r.borrowDetailRepo.BeginTx(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to began tx`, err)
	}
	defer func() {
		defer r.refreshTx(ctx)
		if rec := recover(); rec != nil {
			_ = r.borrowDetailRepo.RollbackTx(ctx)
			panic(r)
		} else if customErr != nil {
			_ = r.borrowDetailRepo.RollbackTx(ctx)
			return
		} else {
			if err := r.borrowDetailRepo.CommitTx(ctx); err != nil {
				log.Println("failed to commit transaction")
				return
			}
			return
		}
	}()
	r.setTx(ctx)

	borrow, err := r.borrowRepo.FindByReferenceID(ctx, request.BorrowID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow record`, err)
	}
	result = &responses.ReturnBooksResponse{
		BorrowID: borrow.ID,
		UserID:   userId,
	}
	for _, bookId := range request.BookIDs {
		_, err := r.borrowDetailRepo.UpdateReturnDateByBorrowIDAndBookID(ctx, borrow.ID, bookId)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get update detail`, err)
		}
		borrowDetail, err := r.borrowDetailRepo.FindByBorrowIDAndBookID(ctx, borrow.ID, bookId)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow detail`, err)
		}
		err = r.bookRepo.UpdateBookStatus(ctx, bookId, enums.Available)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to update book status`, err)
		}
		result.ReturnedBooks = append(result.ReturnedBooks, &responses.ReturnBookResponse{
			BookID:             borrowDetail.BookID,
			VerifiedReturnDate: borrowDetail.VerifiedReturnDate,
		})
	}

	return result, nil
}

var _ usecases_interfaces.ReturnUsecase = &returnUsecase{}

func newReturnUsecase(
	borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	bookRepo repositories_interfaces.BookRepository,
	userRepo repositories_interfaces.UserRepository) *returnUsecase {
	return &returnUsecase{
		borrowRepo:       borrowRepo,
		borrowDetailRepo: borrowDetailRepo,
		bookRepo:         bookRepo,
		userRepo:         userRepo,
	}
}
