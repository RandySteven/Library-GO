package queries

const (
	InsertGenreQuery GoQuery = `
		INSERT INTO genres (genre)
		VALUES (?)
	`

	SelectGenresQuery GoQuery = `
		SELECT id, genre, created_at, updated_at, deleted_at FROM genres
	`

	SelectGenreByID GoQuery = `
		SELECT id, genre FROM genres
		WHERE id = ?
	`
)
