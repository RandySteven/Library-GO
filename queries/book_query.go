package queries

const (
	InsertBookQuery GoQuery = `
		INSERT INTO books (title, description, image, status)
		VALUES (?, ?, ?, ?)
	`

	SelectBooksQuery GoQuery = `
		SELECT id, title, description, image, status, pdf_file, created_at, updated_at, deleted_at 
		FROM
		    books
	`

	SelectBookByIDQuery GoQuery = `
		SELECT id, title, description, image, status, pdf_file, created_at, updated_at, deleted_at 
		FROM
		    books
		WHERE id = ?
	`

	SelectBookAndStatus GoQuery = `
		SELECT id, title, description, image, status, pdf_file, created_at, updated_at, deleted_at
		FROM books
		WHERE id = ? AND status = ?
	`

	UpdateBookStatusQuery GoQuery = `
		UPDATE books
		SET status = ?
		WHERE id = ?
	`

	SelectBookPaginateQuery GoQuery = `
		SELECT id, title, description, image, status, pdf_file, created_at, updated_at, deleted_at
		FROM
		    books
		ORDER BY id
		LIMIT ?
		OFFSET ?
	`
)
