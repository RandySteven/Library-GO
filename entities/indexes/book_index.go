package indexes

import "time"

type BookIndex struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Status      string     `json:"status"`
	Genres      []string   `json:"genres"`
	Authors     []string   `json:"authors"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
