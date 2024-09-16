package requests

type (
	BorrowRequest struct {
		BookIDs []uint64 `json:"book_ids"`
	}
)
