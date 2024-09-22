package queries

const (
	InsertAuthorBookQuery GoQuery = `
		INSERT INTO author_books (author_id, book_id) VALUES (?, ?)
	`
)
