package queries

const (
	InsertBorrowQuery = `INSERT INTO borrows (user_id, borrow_reference) VALUES (?, ?)`
)
