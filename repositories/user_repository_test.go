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
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	email := "test@example.com"
	expectedUser := &models.User{
		ID:          1,
		Name:        "Test User",
		Email:       email,
		PhoneNumber: "123456789",
	}

	// Mock the query result
	rows := sqlmock.NewRows([]string{"id", "name", "address", "email", "phone", "password", "dob", "created_at", "updated_at", "deleted_at", "verified_at"}).
		AddRow(expectedUser.ID, expectedUser.Name, "", expectedUser.Email, expectedUser.PhoneNumber, "", time.Now(), time.Now(), nil, nil, nil)
	mock.ExpectQuery("SELECT .* FROM users WHERE email = ?").WithArgs(email).WillReturnRows(rows)

	repo := repositories.NewRepositories(db)
	user, err := repo.UserRepo.FindByEmail(context.Background(), email)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSave(t *testing.T) {
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
	}

	// Mock the insert query result
	mock.ExpectExec("INSERT INTO users").WithArgs(
		user.Name, user.Address, user.Email, user.PhoneNumber, user.Password, sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositories.NewRepositories(db)
	savedUser, err := repo.UserRepo.Save(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), savedUser.ID)
	assert.Equal(t, user.Email, savedUser.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := uint64(1)
	expectedUser := &models.User{
		ID:          id,
		Name:        "Existing User",
		Email:       "existing@example.com",
		PhoneNumber: "123456789",
	}

	// Mock the query result
	rows := sqlmock.NewRows([]string{"id", "name", "address", "email", "phone", "password", "dob", "created_at", "updated_at", "deleted_at", "verified_at"}).
		AddRow(expectedUser.ID, expectedUser.Name, "", expectedUser.Email, expectedUser.PhoneNumber, "", time.Now(), time.Now(), nil, nil, nil)
	mock.ExpectQuery("SELECT .* FROM users WHERE id = ?").WithArgs(id).WillReturnRows(rows)

	repo := repositories.NewRepositories(db)
	user, err := repo.UserRepo.FindByID(context.Background(), id)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBeginTx(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	repo := repositories.NewRepositories(db)
	err = repo.UserRepo.BeginTx(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, repo.UserRepo.GetTx(context.Background()))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCommitTx(t *testing.T) {
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
