package queries

const (
	InsertAuthorBookQuery GoQuery = `
		INSERT INTO author_books (author_id, book_id) VALUES (?, ?)
	`

	SelectAuthorBookByBookIDQuery GoQuery = `
		SELECT 
		    ab.id, ab.author_id, ab.book_id, ab.created_at, ab.updated_at, ab.deleted_at,
			b.id, b.title, b.description, b.image, b.status, b.created_at, b.updated_at, b.deleted_at,
			a.id, a.name, a.nationality, a.created_at, a.updated_at, a.deleted_at
		FROM author_books AS ab 
			INNER JOIN
		    books AS b 
		ON ab.book_id = b.id
			INNER JOIN
		    authors AS a 
		ON ab.author_id = a.id
		WHERE ab.book_id = ?
	`
)
