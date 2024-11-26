package requests

type (
	RatingRequest struct {
		BookID uint64  `json:"book_id"`
		Rating float32 `json:"rating"`
	}

	RatingFilter struct {
		Order string `json:"order"`
		Limit uint64 `json:"limit"`
	}
)
