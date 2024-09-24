package queries

const (
	InsertBorrowDetailQuery GoQuery = `
		INSERT INTO borrow_details (borrow_id, book_id)
		VALUES 
		    (?, ?)
	`
)
