package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type dashboardUsecase struct {
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository
	bookRepo         repositories_interfaces.BookRepository
	userRepo         repositories_interfaces.UserRepository
}

func (d *dashboardUsecase) SeeAllBorrowRecords(ctx context.Context, pagination *requests.PaginationRequest) (
	result []*responses.ListBorrowDashboardResponse, customErr *apperror.CustomError) {
	borrowDetails, err := d.borrowDetailRepo.FindAll(ctx, 0, 0)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow details`, err)
	}

	for _, borrowDetail := range borrowDetails {
		borrow, err := d.borrowRepo.FindByID(ctx, borrowDetail.BorrowID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get borrow`, err)
		}

		user, err := d.userRepo.FindByID(ctx, borrow.UserID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
		}

		book, err := d.bookRepo.FindByID(ctx, borrowDetail.BookID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get book`, err)
		}

		response := &responses.ListBorrowDashboardResponse{
			ID: borrowDetail.ID,
			User: &struct {
				ID          uint64 `json:"id"`
				Name        string `json:"name"`
				Email       string `json:"email"`
				PhoneNumber string `json:"phone_number"`
			}{
				ID:          user.ID,
				Name:        user.Name,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
			},
			Book: &struct {
				ID     uint64 `json:"id"`
				Title  string `json:"title"`
				Status string `json:"status"`
			}{
				ID:     book.ID,
				Title:  book.Title,
				Status: book.Status.ToString(),
			},
			BorrowReference: borrow.BorrowReference,
			BorrowDate:      borrowDetail.BorrowedDate,
			ReturnDate:      borrowDetail.ReturnedDate,
		}

		result = append(result, response)
	}

	return result, nil
}

var _ usecases_interfaces.DashboardUsecase = &dashboardUsecase{}

func newDashboardUsecase(
	borrowRepo repositories_interfaces.BorrowRepository,
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository,
	bookRepo repositories_interfaces.BookRepository,
	userRepo repositories_interfaces.UserRepository,
) *dashboardUsecase {
	return &dashboardUsecase{
		borrowRepo:       borrowRepo,
		bookRepo:         bookRepo,
		borrowDetailRepo: borrowDetailRepo,
		userRepo:         userRepo,
	}
}
