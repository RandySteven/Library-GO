package usecases_interfaces

import (
	"context"
	"github.com/RandySteven/Library-GO/apperror"
	"github.com/RandySteven/Library-GO/entities/payloads/requests"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	"mime/multipart"
)

type BookUsecase interface {
	AddNewBook(ctx context.Context, request *requests.CreateBookRequest, fileHeader *multipart.FileHeader) (result *responses.CreateBookResponse, customErr *apperror.CustomError)
	GetAllBooks(ctx context.Context) (result []*responses.ListBooksResponse, customErr *apperror.CustomError)
	GetBookByID(ctx context.Context, id uint64) (result *responses.BookDetailResponse, customErr *apperror.CustomError)
	SearchBook(ctx context.Context, request *requests.SearchBookRequest) (result []*responses.ListBooksResponse, customErr *apperror.CustomError)
}
