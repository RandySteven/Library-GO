package queries

const (
	InsertBookGenreQuery GoQuery = `
		INSERT INTO book_genres (book_id, genre_id)
		VALUES (?, ?)
	`

	SelectBookGenreByBookIDQuery GoQuery = `
		SELECT id, book_id, genre_id, created_at, updated_at, deleted_at FROM book_genres WHERE book_id = ?
	`

	SelectBookGenreByGenreIDQuery GoQuery = `
		SELECT id, book_id, genre_id, created_at, updated_at, deleted_at FROM book_genres WHERE genre_id = ?
	`
)
