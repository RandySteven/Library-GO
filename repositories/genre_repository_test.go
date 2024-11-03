package repositories_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/repositories"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type GenreRepositoryTestSuite struct {
	suite.Suite
	ctx  context.Context
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo *repositories.Repositories
}

func (suite *GenreRepositoryTestSuite) SetupSuite() {
	log.Println("Setup suite")
	suite.ctx = context.Background()
	db, mock, err := sqlmock.New()
	suite.NoError(err)
	suite.db = db
	suite.mock = mock
	suite.repo = repositories.NewRepositories(db)
}

func (suite *GenreRepositoryTestSuite) TearDownSuite() {
	log.Println(">> tear down suite")
}

func TestGenreRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GenreRepositoryTestSuite))
}

func (suite *GenreRepositoryTestSuite) TestSave() {
	suite.Run("success to save genre", func() {

		genre := &models.Genre{
			Genre: "genre",
		}
		suite.mock.ExpectPrepare("^INSERT INTO genres \\(genre\\) VALUES \\(\\?\\)$").
			ExpectExec().
			WithArgs(genre.Genre).
			WillReturnResult(sqlmock.NewResult(1, 1))

		savedGenre, err := suite.repo.GenreRepo.Save(context.Background(), genre)
		suite.NoError(err)
		suite.Equal(uint64(1), savedGenre.ID)
		suite.Equal(genre.Genre, savedGenre.Genre)
	})
}

func (suite *GenreRepositoryTestSuite) TestFindByID() {
	suite.Run("success test find by id", func() {
		id := uint64(1)

		expectedGenre := &models.Genre{
			ID:        id,
			Genre:     "genre",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		rows := sqlmock.NewRows([]string{"id", "genre", "created_at", "updated_at", "deleted_at"}).
			AddRow(expectedGenre.ID, expectedGenre.Genre, expectedGenre.CreatedAt, expectedGenre.UpdatedAt, nil)

		suite.mock.ExpectPrepare(queries.SelectGenreByID.ToString()).
			ExpectQuery().
			WithArgs(id).
			WillReturnRows(rows)

		actualGenre, err := suite.repo.GenreRepo.FindByID(suite.ctx, id)
		suite.NoError(err)

		suite.Equal(expectedGenre.ID, actualGenre.ID)
		suite.Equal(expectedGenre.Genre, actualGenre.Genre)
	})
}
