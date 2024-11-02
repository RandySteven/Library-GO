package usecases_test

import (
	"context"
	mocks "github.com/RandySteven/Library-GO/mocks/interfaces/repositories"
	"github.com/stretchr/testify/mock"
	"testing"
)

func InitRepo() {}

func TestGetGenreDetail(t *testing.T) {
	t.Run("success get genre detail", func(t *testing.T) {
		genreRepo := mocks.GenreRepository{}
		bookRepo := mocks.BookRepository{}
		bookGenreRepo := mocks.BookGenreRepository{}
		ctx := context.Background()

		genreRepo.On("FindByID", ctx, 0).Once()
		bookGenreRepo.On("FindBookGenreByBookID", ctx, 0).Once()
		bookRepo.On("FindSelectedBooksId", ctx, mock.AnythingOfType("[]uint64")).Once()

	})
}
