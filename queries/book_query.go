package queries

const (
	InsertBookQuery GoQuery = `
		INSERT INTO books (title, description, image)
		VALUES (?, ?, ?)
	`
)
