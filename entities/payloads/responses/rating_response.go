package responses

type (
	RatingResponse struct {
		ID   uint64 `json:"id"`
		Book struct {
			ID uint64 `json:"id"`
		} `json:"book"`
		User struct {
			ID uint64 `json:"id"`
		} `json:"user"`
		Rating float32 `json:"rating"`
	}

	SortedBookRatings struct {
		Rating float32 `json:"rating"`
		Book   struct {
			ID    uint64 `json:"id"`
			Title string `json:"title"`
			Image string `json:"image"`
		} `json:"book"`
	}
)
