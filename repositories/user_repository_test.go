package repositories_test

import (
	"context"
	_ "database/sql"
	_ "errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/repositories"
	"github.com/stretchr/testify/assert"
)

func TestFindByEmail(t *testing.T) {
	t.Run("success find user by email", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
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

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failed to search user by email", func(t *testing.T) {

	})
}

func TestFindByPhoneNumber(t *testing.T) {
	t.Run("success find user by email", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
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

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failed to search user by email", func(t *testing.T) {

	})
}

func TestSave(t *testing.T) {
	t.Run("success to save user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
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

		assert.NoError(t, err)
		assert.Equal(t, uint64(1), savedUser.ID)
		assert.Equal(t, user.Email, savedUser.Email)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestFindByID(t *testing.T) {
	t.Run("success find user by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
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

		assert.NoError(t, err)

		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.Name, user.Name)
		assert.Equal(t, expectedUser.Address, user.Address)
		assert.Equal(t, expectedUser.Email, user.Email)
		assert.Equal(t, expectedUser.PhoneNumber, user.PhoneNumber)
		assert.Equal(t, expectedUser.Password, user.Password)

		// Use WithinDuration to allow for minor differences in time fields
		assert.WithinDuration(t, expectedUser.DoB, user.DoB, time.Second)
		assert.WithinDuration(t, expectedUser.CreatedAt, user.CreatedAt, time.Second)
		assert.WithinDuration(t, expectedUser.UpdatedAt, user.UpdatedAt, time.Second)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestBeginTx(t *testing.T) {
	t.Run("success test begin tx", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectBegin()
		repo := repositories.NewRepositories(db)
		err = repo.UserRepo.BeginTx(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, repo.UserRepo.GetTx(context.Background()))
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestSetTx(t *testing.T) {
	t.Run("success to set tx", func(t *testing.T) {
		ctx := context.Background()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		tx, _ := db.BeginTx(ctx, nil)

		repo := repositories.NewRepositories(db)
		repo.UserRepo.SetTx(tx)

		assert.NoError(t, err)
		assert.Equal(t, repo.UserRepo.GetTx(ctx), tx)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCommitTx(t *testing.T) {
	t.Run("success test commit tx", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectCommit()

		repo := repositories.NewRepositories(db)
		_ = repo.UserRepo.BeginTx(context.Background())
		err = repo.UserRepo.CommitTx(context.Background())
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRollbackTx(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectRollback()

	repo := repositories.NewRepositories(db)
	_ = repo.UserRepo.BeginTx(context.Background())
	err = repo.UserRepo.RollbackTx(context.Background())
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
