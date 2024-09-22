package responses

import "time"

type (
	CreateBookResponse struct {
		ID string `json:"id"`
	}

	ListBooksResponse struct {
		ID      uint64   `json:"id"`
		Image   string   `json:"image"`
		Title   string   `json:"title"`
		Status  string   `json:"status"`
		Authors []string `json:"authors"`
		Genres  []string `json:"genres"`
	}

	BookDetailResponse struct {
		ID          uint64    `json:"id"`
		Image       string    `json:"image"`
		Title       string    `json:"title"`
		Status      string    `json:"status"`
		Description string    `json:"description"`
		Authors     []string  `json:"authors"`
		Genres      []string  `json:"genres"`
		CreatedAt   time.Time `json:"created_at"`
	}
)
