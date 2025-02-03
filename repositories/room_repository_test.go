package repositories_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/entities/models"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	"github.com/RandySteven/Library-GO/repositories"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RoomRepositoryTestSuite struct {
	ctx      context.Context
	tx       *sql.Tx
	db       *sql.DB
	roomRepo repositories_interfaces.RoomRepository
	mock     sqlmock.Sqlmock
	suite.Suite
}

func (suite *RoomRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	suite.NoError(err)
	suite.db = db
	suite.mock = mock
	suite.roomRepo = repositories.NewRepositories(db).RoomRepo
}

func TestRoomRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomRepositoryTestSuite))
}

func (suite *RoomRepositoryTestSuite) TestSave() {
	suite.Run("success to save room", func() {
		room := &models.Room{
			Name:      "Room A",
			Thumbnail: "",
		}

		queries := `INSERT INTO rooms (name, thumbnail, is_available) VALUES (?, ?, true)`

		suite.mock.ExpectPrepare(queries).
			ExpectExec().
			WithArgs(&room.Name, &room.Thumbnail).
			WillReturnResult(
				sqlmock.NewResult(1, 1),
			)

		result, err := suite.roomRepo.Save(suite.ctx, room)
		suite.NoError(err)
		suite.Equal(uint64(0x1), result.ID)
		suite.Equal(room.Name, result.Name)
	})

	suite.Run("failed to save room", func() {

	})
}
