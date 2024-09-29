package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type (
	BorrowUsecase interface {
		BorrowTransaction(ctx context.Context) (result *responses.BorrowResponse, customErr *apperror.CustomError)
		GetAllBorrows(ctx context.Context) (result []*responses.BorrowListResponse, customErr *apperror.CustomError)
		GetBorrowDetail(ctx context.Context, id uint64) (result *responses.BorrowDetailResponse, customErr *apperror.CustomError)
	}
)
