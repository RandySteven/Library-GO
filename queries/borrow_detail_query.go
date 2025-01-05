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

	SelectBorrowDetailReturnedDateTodayQuery GoQuery = `
		SELECT id, borrow_id, book_id, borrowed_date, returned_date, verified_return_date, created_at, updated_at, deleted_at
		FROM borrow_details
		WHERE returned_date IS CURRENT_DATE() AND verified_return_date IS NULL
	`

	SelectBorrowDetailWithBookIDQuery GoQuery = `
		SELECT id, borrow_id, book_id, borrowed_date, returned_date, verified_return_date, created_at, updated_at, deleted_at
		FROM borrow_details
		WHERE book_id = ?
	`

	SelectBorrowDetailByBorrowAndBookQuery GoQuery = `
		SELECT id, borrow_id, book_id, borrowed_date, verified_return_date, verified_return_date, created_at, updated_at, deleted_at
		FROM borrow_details
		WHERE borrow_id = ? AND book_id = ?
	`

	UpdateBorrowReturnDateByBorrowAndBookQuery GoQuery = `
		UPDATE borrow_details 
		    SET verified_return_date = NOW(),
		                               updated_at = NOW()
		WHERE borrow_id = ? AND book_id = ?
	`
)
