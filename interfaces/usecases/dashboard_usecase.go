package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
)

type DashboardUsecase interface {
	SeeAllBorrowRecords(ctx context.Context, pagination *requests.PaginationRequest) (result []*responses.ListBorrowDashboardResponse, customErr *apperror.CustomError)
}
