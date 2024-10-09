package requests

type (
	BagRequest struct {
		BookID uint64 `json:"book_id"`
	}

	DeleteBookBagRequest struct {
		BookIDs []uint64 `json:"book_ids"`
	}
)
