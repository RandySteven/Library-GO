package responses

type (
	AddBagResponse struct {
		BookID uint64 `json:"book_id"`
	}

	BookBagResponse struct {
		ID    uint64 `json:"id"`
		Title string `json:"title"`
		Image string `json:"image"`
	}

	GetAllBagsResponse struct {
		Books []*BookBagResponse `json:"books"`
	}
)
