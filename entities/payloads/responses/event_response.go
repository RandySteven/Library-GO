package responses

import "time"

type (
	EventCreateResponse struct {
		ID        string     `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	ListEventResponse struct {
		ID        uint64     `json:"id"`
		Title     string     `json:"title"`
		Thumbnail string     `json:"thumbnail"`
		Price     *uint64    `json:"price"`
		Slot      uint64     `json:"slot"`
		Date      time.Time  `json:"date"`
		StartTime time.Time  `json:"start_time"`
		EndTime   time.Time  `json:"end_time"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	EventDetailResponse struct {
		ID          uint64     `json:"id"`
		Title       string     `json:"title"`
		Thumbnail   string     `json:"thumbnail"`
		Description string     `json:"description"`
		Price       *uint64    `json:"price"`
		Slot        uint64     `json:"slot"`
		Date        time.Time  `json:"date"`
		StartTime   time.Time  `json:"start_time"`
		EndTime     time.Time  `json:"end_time"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	}
)
