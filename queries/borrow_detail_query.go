package queries

const (
	InsertBorrowDetailQuery GoQuery = `
		INSERT INTO borrow_details (borrow_id, book_id, returned_date)
		VALUES 
		    (?, ?, ?)
	`

	SelectBorrowDetailByBorrowIDQuery GoQuery = `
		SELECT id, borrow_id, book_id, borrowed_date, returned_date, verified_return_date, created_at, updated_at, deleted_at
		FROM borrow_details
		WHERE borrow_id = ?
	`
)
