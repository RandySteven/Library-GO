package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type returnUsecase struct {
	borrowRepo       repositories_interfaces.BorrowRepository
	borrowDetailRepo repositories_interfaces.BorrowDetailRepository
	bookRepo         repositories_interfaces.BookRepository
	userRepo         repositories_interfaces.UserRepository
}

func (r *returnUsecase) ReturnBook(ctx context.Context, request *requests.ReturnRequest) (result *responses.ReturnBooksResponse, customErr *apperror.CustomError) {
	return
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
