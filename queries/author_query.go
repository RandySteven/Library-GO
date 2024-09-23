package queries

const (
	InsertAuthorQuery GoQuery = `
		INSERT INTO authors (name, nationality)
		VALUES
		    (?, ?)
	`

	SelectAuthorQuery GoQuery = `
		SELECT id, name, nationality, created_at, updated_at, deleted_at FROM authors
	`

	SelectAuthorByID GoQuery = `
		SELECT id, name, nationality, created_at, updated_at, deleted_at FROM authors
		WHERE id = ?
	`
)
