package repositories

import (
	"database/sql"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type Repositories struct {
	UserRepo         repositories_interfaces.UserRepository
	BookRepo         repositories_interfaces.BookRepository
	BorrowRepo       repositories_interfaces.BorrowRepository
	BorrowDetailRepo repositories_interfaces.BorrowDetailRepository
	GenreRepo        repositories_interfaces.GenreRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepo: newUserRepository(db),
	}
}
