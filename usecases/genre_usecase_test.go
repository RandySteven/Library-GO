package usecases_test

import (
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	mocks "github.com/RandySteven/Library-GO/mocks/interfaces/repositories"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenreUsecaseTestSuite struct {
	suite.Suite
	genreRepo     *mocks.GenreRepository
	bookRepo      *mocks.BookRepository
	bookGenreRepo *mocks.BookGenreRepository
	usecase       usecases_interfaces.GenreUsecase
}

func (suite *GenreUsecaseTestSuite) SetupSuite() {
	suite.genreRepo = new(mocks.GenreRepository)
	suite.bookRepo = new(mocks.BookRepository)
	suite.bookGenreRepo = new(mocks.BookGenreRepository)
}

func TestGenreUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(GenreUsecaseTestSuite))
}

func (suite *GenreUsecaseTestSuite) TestAddGenre() {
	suite.Run("success to save", func() {

	})
}
