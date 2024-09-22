package queries

const (
	InsertBookGenreQuery GoQuery = `
		INSERT INTO book_genres (book_id, genre_id)
		VALUES (?, ?)
	`
)
