package repositories

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookRepositoryTestSuite struct {
	suite.Suite
	db   *sql.DB
	tx   *sql.Tx
	mock sqlmock.Sqlmock
	ctx  context.Context
	repo *bookRepository
}

func (suite *BookRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, mock, err := sqlmock.New()
	suite.NoError(err)
	suite.db = db
	suite.mock = mock
	suite.repo = newBookRepository(suite.db)
}

func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}

func (suite *BookRepositoryTestSuite) TestSave() {
	suite.Run("success to save book", func() {
		ctx := context.Background()
		book := &models.Book{
			Title:       "Test Book",
			Description: "Test description on book",
			Image:       "test_image",
			Status:      enums.Available,
		}
		//result, err := suite.repo.Save(ctx, book)
		result, err := suite.db.ExecContext(ctx, `INSERT INTO books (title, description, image, status) VALUES (?, ?, ?, ?)`, book.Title, book.Description, book.Image, book.Status)
		suite.NoError(err)
		id, err := result.LastInsertId()
		suite.NotEqual(id, 0)
		suite.NoError(err)
	})
}
