package responses

import "time"

type (
	GenreResponseDetail struct {
		ID        uint64     `json:"id"`
		Genre     string     `json:"genre"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	ListGenresResponse struct {
		ID    uint64 `json:"id"`
		Genre string `json:"genre"`
	}
)
