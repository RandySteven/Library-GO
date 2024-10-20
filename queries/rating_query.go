package queries

const (
	InsertIntoRatingQuery GoQuery = `
		INSERT INTO ratings (book_id, user_id, score)
		VALUES 
		    (?, ?, ?)
	`

	SelectRatingsQuery GoQuery = `
		SELECT id, book_id, user_id, score, created_at, updated_at, deleted_at
		FROM
		    ratings
	`

	SelectRatingValue GoQuery = `
		SELECT book_id, AVG(score) as score
		FROM
		    ratings
		WHERE book_id = ?
		GROUP BY book_id
	`
)
