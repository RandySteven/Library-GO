package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type (
	BorrowUsecase interface {
		BorrowTransaction(ctx context.Context) (result *responses.BorrowResponse, customErr *apperror.CustomError)
		GetAllBorrows(ctx context.Context) (result []*responses.BorrowListResponse, customErr *apperror.CustomError)
		GetBorrowDetail(ctx context.Context, borrowId string) (result *responses.BorrowDetailResponse, customErr *apperror.CustomError)
		ReturnBorrowBook(ctx context.Context, request *requests.ReturnRequest) (result *responses.ReturnBooksResponse, customErr *apperror.CustomError)
	}
)
