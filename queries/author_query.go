package queries

const (
	InsertAuthorQuery GoQuery = `
		INSERT INTO authors (name, nationality)
		VALUES
		    (?, ?)
	`

	SelectAuthorQuery GoQuery = `
		SELECT id, name, nationality FROM authors
	`

	SelectAuthorByID GoQuery = `
		SELECT id, name, nationality FROM authors
		WHERE id = ?
	`
)
