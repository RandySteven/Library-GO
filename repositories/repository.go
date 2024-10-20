package repositories

import (
	"database/sql"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
)

type Repositories struct {
	AuthorRepo       repositories_interfaces.AuthorRepository
	AuthorBookRepo   repositories_interfaces.AuthorBookRepository
	BagRepo          repositories_interfaces.BagRepository
	BookRepo         repositories_interfaces.BookRepository
	BookGenreRepo    repositories_interfaces.BookGenreRepository
	BorrowRepo       repositories_interfaces.BorrowRepository
	BorrowDetailRepo repositories_interfaces.BorrowDetailRepository
	GenreRepo        repositories_interfaces.GenreRepository
	UserRepo         repositories_interfaces.UserRepository
	RatingRepo       repositories_interfaces.RatingRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		AuthorRepo:       newAuthorRepository(db),
		AuthorBookRepo:   newAuthorBookRepository(db),
		BagRepo:          newBagRepository(db),
		BookRepo:         newBookRepository(db),
		BookGenreRepo:    newBookGenreRepository(db),
		BorrowRepo:       newBorrowRepository(db),
		BorrowDetailRepo: newBorrowDetailRepository(db),
		GenreRepo:        newGenreRepository(db),
		UserRepo:         newUserRepository(db),
		RatingRepo:       newRatingRepository(db),
	}
}
