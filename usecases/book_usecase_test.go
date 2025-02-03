package usecases

import (
	mocks "github.com/RandySteven/Library-GO/mocks/interfaces/repositories"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookUsecaseTestSuite struct {
	suite.Suite
	bookRepoMock mocks.BookRepository
}

func TestBookUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BookUsecaseTestSuite))
}
