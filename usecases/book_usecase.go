package usecases

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
)

type bookUsecase struct {
	userRepo repositories_interfaces.UserRepository
	bookRepo repositories_interfaces.BookRepository
}

func (b *bookUsecase) AddNewBook(ctx context.Context, request *requests.CreateBookRequest) (result *responses.CreateBookResponse, customErr *apperror.CustomError) {
	return
}

func (b *bookUsecase) GetAllBooks(ctx context.Context) (result []*responses.ListBooksResponse, customErr *apperror.CustomError) {
	return
}

func (b *bookUsecase) GetBookByID(ctx context.Context, id uint64) (result *responses.BookDetailResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases_interfaces.BookUsecase = &bookUsecase{}

func newBookUsecase(
	userRepo repositories_interfaces.UserRepository,
	bookRepo repositories_interfaces.BookRepository) *bookUsecase {
	return &bookUsecase{
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}
