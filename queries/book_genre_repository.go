package queries

const (
	InsertBookGenreQuery GoQuery = `
		INSERT INTO book_genres (book_id, genre_id)
		VALUES (?, ?)
	`

	SelectBookGenreByBookIDQuery GoQuery = `
		SELECT bg.id, bg.book_id, bg.genre_id, bg.created_at, bg.updated_at, bg.deleted_at,
		       b.id, b.title, b.description, b.image, b.status, b.created_at, b.updated_at, b.deleted_at,
		       g.id, g.genre, g.created_at, g.updated_at, g.deleted_at
		FROM book_genres AS bg
				 INNER JOIN books as b
							ON bg.book_id = b.id
				 INNER JOIN genres as g
							ON bg.genre_id = g.id
			WHERE bg.book_id = ?
	`

	SelectBookGenreByGenreIDQuery GoQuery = `
		SELECT bg.id, bg.book_id, bg.genre_id, bg.created_at, bg.updated_at, bg.deleted_at,
		       b.id, b.title, b.description, b.image, b.status, b.created_at, b.updated_at, b.deleted_at,
		       g.id, g.genre, g.created_at, g.updated_at, g.deleted_at
		FROM book_genres AS bg
				 INNER JOIN books as b
							ON bg.book_id = b.id
				 INNER JOIN genres as g
							ON bg.genre_id = g.id
		WHERE bg.genre_id = ?
	`
)
