package repositories_test

import (
	"context"
	_ "database/sql"
	_ "errors"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/repositories"
)

type UserRepositoryTestSuite struct {
	suite.Suite
}

func (u *UserRepositoryTestSuite) SetupSuite() {
	log.Println("Setup suite")
}

func (u *UserRepositoryTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (u *UserRepositoryTestSuite) TestFindByEmail() {
	u.Run("success find user by email", func() {
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		email := "test@example.com"
		expectedUser := &models.User{
			ID:          1,
			Name:        "Test User",
			Email:       email,
			PhoneNumber: "123456789",
			DoB:         time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Mock the query result
		rows := sqlmock.NewRows([]string{"id", "name", "address", "email", "phone", "password", "dob", "created_at", "updated_at", "deleted_at", "verified_at"}).
			AddRow(expectedUser.ID, expectedUser.Name, "", expectedUser.Email, expectedUser.PhoneNumber, "", expectedUser.DoB, expectedUser.CreatedAt, expectedUser.UpdatedAt, nil, nil)
		mock.ExpectQuery("SELECT .* FROM users WHERE email = ?").WithArgs(email).WillReturnRows(rows)

		repo := repositories.NewRepositories(db)
		user, err := repo.UserRepo.FindByEmail(context.Background(), email)

		u.NoError(err)
		u.Equal(expectedUser, user)
		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestFindByPhoneNumber() {
	u.Run("success find user by email", func() {
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		email := "test@example.com"
		phone := "123456789"
		expectedUser := &models.User{
			ID:          1,
			Name:        "Test User",
			Email:       email,
			PhoneNumber: phone,
			DoB:         time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Mock the query result
		rows := sqlmock.NewRows([]string{"id", "name", "address", "email", "phone_number", "password", "dob", "created_at", "updated_at", "deleted_at", "verified_at"}).
			AddRow(expectedUser.ID, expectedUser.Name, "", expectedUser.Email, expectedUser.PhoneNumber, "", expectedUser.DoB, expectedUser.CreatedAt, expectedUser.UpdatedAt, nil, nil)
		mock.ExpectQuery("SELECT .* FROM users WHERE phone_number = ?").WithArgs(phone).WillReturnRows(rows)

		repo := repositories.NewRepositories(db)
		user, err := repo.UserRepo.FindByPhoneNumber(context.Background(), phone)

		u.NoError(err)
		u.Equal(expectedUser, user)
		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestSave() {
	u.Run("success to save user", func() {
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		// Create a user to be saved
		user := &models.User{
			Name:        "New User",
			Address:     "123 Street",
			Email:       "newuser@example.com",
			PhoneNumber: "987654321",
			Password:    "password",
			DoB:         time.Now(),
		}

		// Mock the insert query result
		mock.ExpectPrepare("^INSERT INTO users \\(name, address, email, phone_number, password, dob\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)$").
			ExpectExec().
			WithArgs(user.Name, user.Address, user.Email, user.PhoneNumber, user.Password, user.DoB).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := repositories.NewRepositories(db)
		savedUser, err := repo.UserRepo.Save(context.Background(), user)

		u.NoError(err)
		u.Equal(uint64(1), savedUser.ID)
		u.Equal(user.Email, savedUser.Email)
		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestFindByID() {
	u.Run("success find user by id", func() {
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		id := uint64(1)
		now := time.Now()

		// Define the expected user
		expectedUser := &models.User{
			ID:          id,
			Name:        "Existing User",
			Address:     "123 street",
			Email:       "existing@example.com",
			Password:    "test_1234",
			PhoneNumber: "123456789",
			DoB:         now,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Mock the database response
		rows := sqlmock.NewRows([]string{"id", "name", "address", "email", "phone_number", "password", "dob", "created_at", "updated_at", "deleted_at", "verified_at"}).
			AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Address, expectedUser.Email, expectedUser.PhoneNumber, expectedUser.Password, expectedUser.DoB, expectedUser.CreatedAt, expectedUser.UpdatedAt, nil, nil)

		// Ensure the query is matched correctly
		mock.ExpectPrepare(`SELECT id, name, address, email, phone_number, password, dob, created_at, updated_at, deleted_at, verified_at FROM users WHERE id = ?`).
			ExpectQuery().
			WithArgs(id).
			WillReturnRows(rows)

		repo := repositories.NewRepositories(db)
		user, err := repo.UserRepo.FindByID(context.Background(), id)

		u.NoError(err)

		u.Equal(expectedUser.ID, user.ID)
		u.Equal(expectedUser.Name, user.Name)
		u.Equal(expectedUser.Address, user.Address)
		u.Equal(expectedUser.Email, user.Email)
		u.Equal(expectedUser.PhoneNumber, user.PhoneNumber)
		u.Equal(expectedUser.Password, user.Password)

		// Use WithinDuration to allow for minor differences in time fields
		u.WithinDuration(expectedUser.DoB, user.DoB, time.Second)
		u.WithinDuration(expectedUser.CreatedAt, user.CreatedAt, time.Second)
		u.WithinDuration(expectedUser.UpdatedAt, user.UpdatedAt, time.Second)

		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestBeginTx() {
	u.Run("success test begin tx", func() {
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		mock.ExpectBegin()
		repo := repositories.NewRepositories(db)
		err = repo.UserRepo.BeginTx(context.Background())
		u.NoError(err)
		u.NotNil(repo.UserRepo.GetTx(context.Background()))
		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestSetTx() {
	u.Run("success to set tx", func() {
		ctx := context.Background()
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()
		tx, _ := db.BeginTx(ctx, nil)

		repo := repositories.NewRepositories(db)
		repo.UserRepo.SetTx(tx)

		u.NoError(err)
		u.Equal(repo.UserRepo.GetTx(ctx), tx)
		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestCommitTx() {
	u.Run("success test commit tx", func() {
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectCommit()

		repo := repositories.NewRepositories(db)
		_ = repo.UserRepo.BeginTx(context.Background())
		err = repo.UserRepo.CommitTx(context.Background())
		u.NoError(err)
		u.NoError(mock.ExpectationsWereMet())
	})
}

func (u *UserRepositoryTestSuite) TestRollbackTx() {
	db, mock, err := sqlmock.New()
	u.NoError(err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectRollback()

	repo := repositories.NewRepositories(db)
	_ = repo.UserRepo.BeginTx(context.Background())
	err = repo.UserRepo.RollbackTx(context.Background())
	u.NoError(err)
	u.NoError(mock.ExpectationsWereMet())
}

func (u *UserRepositoryTestSuite) TestInitTrigger() {
	u.Run("success return trigger by db", func() {
		ctx := context.Background()
		db, _, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()

		repo := repositories.NewRepositories(db)
		repo.UserRepo.InitTrigger()

		u.Nil(repo.UserRepo.GetTx(ctx))
	})

	u.Run("success return trigger by tx", func() {
		ctx := context.Background()
		db, mock, err := sqlmock.New()
		u.NoError(err)
		defer db.Close()
		mock.ExpectBegin()

		repo := repositories.NewRepositories(db)
		repo.UserRepo.BeginTx(ctx)
		repo.UserRepo.InitTrigger()

		u.NotNil(repo.UserRepo.GetTx(ctx))
	})
}
