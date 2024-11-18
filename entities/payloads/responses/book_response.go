package responses

import "time"

type (
	CreateBookResponse struct {
		ID string `json:"id"`
	}

	ListBooksResponse struct {
		ID        uint64     `json:"id"`
		Image     string     `json:"image"`
		Title     string     `json:"title"`
		Status    string     `json:"status"`
		Rating    float32    `json:"rating"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	AuthorBookResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	GenreBookResponse struct {
		ID    uint64 `json:"id"`
		Genre string `json:"genre"`
	}

	BookDetailResponse struct {
		ID          uint64                `json:"id"`
		Rating      float32               `json:"rating"`
		Image       string                `json:"image"`
		Title       string                `json:"title"`
		Status      string                `json:"status"`
		Description string                `json:"description"`
		Authors     []*AuthorBookResponse `json:"authors"`
		Genres      []*GenreBookResponse  `json:"genres"`
		CreatedAt   time.Time             `json:"created_at"`
	}
)
