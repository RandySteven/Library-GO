package queries

const (
	InsertBorrowQuery             = `INSERT INTO borrows (user_id, borrow_reference) VALUES (?, ?)`
	SelectBorrowByIDQuery GoQuery = `
		SELECT id, user_id, borrow_reference, created_at, updated_at, deleted_at
		FROM borrows
		WHERE id = ?
	`
	SelectBorrowByReference GoQuery = `
		SELECT id, user_id, borrow_reference, created_at, updated_at, deleted_at
		FROM borrows
		WHERE borrow_reference = ?
	`
	SelectBorrowUserIdQuery GoQuery = `
		SELECT id, user_id, borrow_reference, created_at, updated_at, deleted_at
		FROM borrows
		WHERE user_id = ?
	`

	SelectBorrowsQueryWithUser GoQuery = `
		SELECT b.id, b.user_id, u.name, u.email
		FROM
		    borrows b
		INNER JOIN
			users u
		ON b.user_id = u.id
	`
)
