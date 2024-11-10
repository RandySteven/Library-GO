package queries

const (
	InsertCommentQuery GoQuery = `
		INSERT INTO comments (user_id, book_id, parent_id, reply_id, comment)
		VALUES 
		    (?, ?, ?, ?, ?)
	`

	SelectBookCommentsQuery GoQuery = `
		SELECT id, user_id, book_id, parent_id, comment, created_at, updated_at, deleted_at
		FROM comments
		WHERE book_id = ?
	`

	SelectCommentByIDQuery GoQuery = `
		SELECT id, user_id, book_id, parent_id, comment, created_at, updated_at, deleted_at
		FROM comments
		WHERE id = ?
	`
)
