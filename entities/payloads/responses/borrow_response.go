package responses

import "time"

type (
	BorrowResponse struct {
		ID           uint64    `json:"id"`
		BorrowID     string    `json:"borrow_id"`
		UserID       uint64    `json:"user_id"`
		TotalItems   uint64    `json:"total_items"`
		Status       string    `json:"status"`
		BorrowedDate time.Time `json:"borrowed_date"`
		ReturnedDate time.Time `json:"returned_date"`
		CreatedAt    time.Time `json:"created_at"`
	}

	BorrowListResponse struct {
		ID              uint64    `json:"id"`
		BorrowReference string    `json:"borrow_id"`
		BorrowedDate    time.Time `json:"borrowed_date"`
	}

	BorrowedBook struct {
		ID           uint64    `json:"id"`
		Title        string    `json:"title"`
		Image        string    `json:"image"`
		BorrowedDate time.Time `json:"borrowed_date"`
		ReturnedDate time.Time `json:"returned_date"`
		HasReturned  bool      `json:"has_returned"`
	}

	BorrowDetailResponse struct {
		ID              uint64 `json:"id"`
		BorrowReference string `json:"borrow_reference"`
		TotalItems      int    `json:"total_items"`
		User            struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"user"`
		BorrowedBooks []*BorrowedBook `json:"borrowed_books"`
	}

	ReturnBookResponse struct {
		BookID             uint64     `json:"book_id"`
		VerifiedReturnDate *time.Time `json:"verified_return_date"`
	}

	ReturnBooksResponse struct {
		ID            uint64                `json:"id"`
		BorrowID      uint64                `json:"borrow_id"`
		UserID        uint64                `json:"user_id"`
		ReturnedBooks []*ReturnBookResponse `json:"returned_books"`
	}
)
