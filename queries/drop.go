package queries

const (
	DropUserTable         DropQuery = "DROP TABLE IF EXISTS users"
	DropAuthorTable       DropQuery = "DROP TABLE IF EXISTS authors"
	DropGenreTable        DropQuery = "DROP TABLE IF EXISTS genres"
	DropBookTable         DropQuery = "DROP TABLE IF EXISTS books"
	DropBookGenreTable    DropQuery = "DROP TABLE IF EXISTS book_genres"
	DropAuthorBookTable   DropQuery = "DROP TABLE IF EXISTS author_books"
	DropBagTable          DropQuery = "DROP TABLE IF EXISTS bags"
	DropBorrowTable       DropQuery = "DROP TABLE IF EXISTS borrows"
	DropBorrowDetailTable DropQuery = "DROP TABLE IF EXISTS borrow_details"
	DropRatingTable       DropQuery = "DROP TABLE IF EXISTS ratings"
)
