package queries

const (
	InsertUserGenreQuery GoQuery = `
		INSERT INTO user_genres (user_id, genre_id)
		VALUES 
		    (?, ?)
	`
)
