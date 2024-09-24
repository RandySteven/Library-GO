package queries

const (
	InsertAuthorBookQuery GoQuery = `
		INSERT INTO author_books (author_id, book_id) VALUES (?, ?)
	`

	SelectAuthorBookByBookIDQuery GoQuery = `
		SELECT id, author_id, book_id, created_at, updated_at, deleted_at FROM author_books WHERE book_id = ?
	`
)
