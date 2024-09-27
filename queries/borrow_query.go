package queries

const (
	InsertBorrowQuery             = `INSERT INTO borrows (user_id, borrow_reference) VALUES (?, ?)`
	SelectBorrowByIDQuery GoQuery = `
		SELECT id, user_id, borrow_reference, created_at, updated_at, deleted_at
		FROM borrows
		WHERE id = ?
	`
)
