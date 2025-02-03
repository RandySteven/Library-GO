package repositories_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/enums"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/repositories"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AuthorBookRepositoryTestSuite struct {
	suite.Suite
	db   *sql.DB
	tx   *sql.Tx
	mock sqlmock.Sqlmock
	ctx  context.Context
	repo repositories_interfaces.AuthorBookRepository
}

func (suite *AuthorBookRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	suite.NoError(err)
	suite.db = db
	suite.mock = mock
	suite.repo = repositories.NewRepositories(db).AuthorBookRepo
}

func TestAuthorBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorBookRepositoryTestSuite))
}

func (suite *AuthorBookRepositoryTestSuite) TestSave() {
	suite.Run("success to save author book", func() {
		authorBook := &models.AuthorBook{
			AuthorID: 1,
			BookID:   1,
		}

		suite.mock.ExpectPrepare(
			"INSERT INTO author_books (author_id, book_id) VALUES (?, ?)",
		).
			ExpectExec().
			WithArgs(
				authorBook.AuthorID, authorBook.BookID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		savedAuthorBook, err := suite.repo.Save(suite.ctx, authorBook)
		suite.NoError(err)
		suite.Equal(authorBook.BookID, savedAuthorBook.BookID)
		suite.Equal(authorBook.AuthorID, savedAuthorBook.AuthorID)
		suite.NoError(suite.mock.ExpectationsWereMet())
	})

	suite.Run("failed to save author book", func() {
		authorBook := &models.AuthorBook{
			AuthorID: 1,
			BookID:   1,
		}

		suite.mock.ExpectPrepare(
			"INSERT INTO author_books (author_id, book_id) VALUES (?, ?)",
		).
			ExpectExec().
			WillReturnError(fmt.Errorf(`failed to insert db`))

		savedAuthorBook, err := suite.repo.Save(suite.ctx, authorBook)
		suite.Error(err)
		suite.Nil(savedAuthorBook)
	})
}

func (suite *AuthorBookRepositoryTestSuite) TestFindAuthorBookByBookID() {
	suite.Run("success to find author books by book_id", func() {
		authorBooks := []*models.AuthorBook{
			{
				uint64(1),
				uint64(1),
				uint64(1),
				time.Now(),
				time.Now(),
				nil,
				models.Book{
					uint64(1),
					"Book A",
					"Description of Book A",
					"imageURL book A",
					enums.Available,
					time.Now(),
					time.Now(),
					nil,
				},
				models.Author{
					uint64(1),
					"Luis",
					"Indonesia",
					time.Now(),
					time.Now(),
					nil,
				},
			},
			{uint64(2),
				uint64(1),
				uint64(1),
				time.Now(),
				time.Now(),
				nil,
				models.Book{
					uint64(1),
					"Book A",
					"Description of Book A",
					"imageURL book A",
					enums.Available,
					time.Now(),
					time.Now(),
					nil,
				},
				models.Author{
					uint64(1),
					"Luis",
					"Indonesia",
					time.Now(),
					time.Now(),
					nil,
				}},
			{uint64(3),
				uint64(1),
				uint64(2),
				time.Now(),
				time.Now(),
				nil,
				models.Book{
					uint64(2),
					"Book B",
					"Description of Book B",
					"imageURL book B",
					enums.Available,
					time.Now(),
					time.Now(),
					nil,
				},
				models.Author{
					uint64(1),
					"Luis",
					"Indonesia",
					time.Now(),
					time.Now(),
					nil,
				}},
		}

		rows := sqlmock.NewRows([]string{
			"ab.id", "ab.author_id", "ab.book_id", "ab.created_at", "ab.updated_at", "ab.deleted_at",
			"b.id", "b.title", "b.description", "b.image", "b.status", "b.created_at", "b.updated_at", "b.deleted_at",
			"a.id", "a.name", "a.nationality", "a.created_at", "a.updated_at", "a.deleted_at",
		}).
			AddRow(authorBooks[0].ID, authorBooks[0].AuthorID, authorBooks[0].BookID, authorBooks[0].CreatedAt, authorBooks[0].UpdatedAt, authorBooks[0].DeletedAt,
				authorBooks[0].Book.ID, authorBooks[0].Book.Title, authorBooks[0].Book.Description, authorBooks[0].Book.Image, authorBooks[0].Book.Status, authorBooks[0].Book.CreatedAt, authorBooks[0].Book.UpdatedAt, authorBooks[0].Book.DeletedAt,
				authorBooks[0].Author.ID, authorBooks[0].Author.Name, authorBooks[0].Author.Nationality, authorBooks[0].Author.CreatedAt, authorBooks[0].Author.UpdatedAt, authorBooks[0].Author.DeletedAt).
			AddRow(authorBooks[1].ID, authorBooks[1].AuthorID, authorBooks[1].BookID, authorBooks[1].CreatedAt, authorBooks[1].UpdatedAt, authorBooks[1].DeletedAt,
				authorBooks[1].Book.ID, authorBooks[1].Book.Title, authorBooks[1].Book.Description, authorBooks[1].Book.Image, authorBooks[1].Book.Status, authorBooks[1].Book.CreatedAt, authorBooks[1].Book.UpdatedAt, authorBooks[1].Book.DeletedAt,
				authorBooks[1].Author.ID, authorBooks[1].Author.Name, authorBooks[1].Author.Nationality, authorBooks[1].Author.CreatedAt, authorBooks[1].Author.UpdatedAt, authorBooks[1].Author.DeletedAt).
			AddRow(authorBooks[2].ID, authorBooks[2].AuthorID, authorBooks[2].BookID, authorBooks[2].CreatedAt, authorBooks[2].UpdatedAt, authorBooks[2].DeletedAt,
				authorBooks[2].Book.ID, authorBooks[2].Book.Title, authorBooks[2].Book.Description, authorBooks[2].Book.Image, authorBooks[2].Book.Status, authorBooks[2].Book.CreatedAt, authorBooks[2].Book.UpdatedAt, authorBooks[2].Book.DeletedAt,
				authorBooks[2].Author.ID, authorBooks[2].Author.Name, authorBooks[2].Author.Nationality, authorBooks[2].Author.CreatedAt, authorBooks[2].Author.UpdatedAt, authorBooks[2].Author.DeletedAt)

		bookID := uint64(1)
		suite.mock.ExpectQuery(
			`
					SELECT 
						ab.id, ab.author_id, ab.book_id, ab.created_at, ab.updated_at, ab.deleted_at,
						b.id, b.title, b.description, b.image, b.status, b.created_at, b.updated_at, b.deleted_at,
						a.id, a.name, a.nationality, a.created_at, a.updated_at, a.deleted_at
					FROM author_books AS ab 
						INNER JOIN
						books AS b 
					ON ab.book_id = b.id
						INNER JOIN
						authors AS a 
					ON ab.author_id = a.id
					WHERE ab.book_id = ?
				`,
		).WithArgs(bookID).WillReturnRows(rows)

		result, err := suite.repo.FindAuthorBookByBookID(suite.ctx, bookID)
		suite.NoError(err)
		suite.NotNil(result)
		suite.Equal(2, len(result))
	})
}
