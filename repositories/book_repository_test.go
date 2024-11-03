package repositories_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/repositories"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookRepositoryTestSuite struct {
	suite.Suite
	db   *sql.DB
	tx   *sql.Tx
	mock sqlmock.Sqlmock
	repo *repositories.Repositories
	ctx  context.Context
}

func (suite *BookRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, mock, err := sqlmock.New()
	suite.NoError(err)
	suite.db = db
	suite.mock = mock
}

func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}

func (suite *BookRepositoryTestSuite) TestSave() {
	suite.Run("success to save book", func() {

	})
}
