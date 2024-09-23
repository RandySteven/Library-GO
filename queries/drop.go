package queries

const (
	DropUserTable       DropQuery = "DROP TABLE IF EXISTS users"
	DropAuthorTable     DropQuery = "DROP TABLE IF EXISTS authors"
	DropGenreTable      DropQuery = "DROP TABLE IF EXISTS genres"
	DropBookTable       DropQuery = "DROP TABLE IF EXISTS books"
	DropBookGenreTable  DropQuery = "DROP TABLE IF EXISTS book_genres"
	DropAuthorBookTable DropQuery = "DROP TABLE IF EXISTS author_books"
)
