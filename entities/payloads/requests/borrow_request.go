package requests

type (
	BorrowRequest struct {
		UserID uint64 `json:"user_id"`
	}

	ReturnRequest struct {
		BorrowID string   `json:"borrow_id"`
		BookIDs  []uint64 `json:"book_id"`
	}

	ConfirmBorrowRequest struct {
		BorrowID string `json:"borrow_id"`
	}
)
