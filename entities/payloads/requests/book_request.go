package requests

type (
	CreateBookRequest struct {
		Title       string   `json:"title" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Genres      []uint64 `json:"genres" validate:"required"`
		Authors     []uint64 `json:"authors" validate:"required"`
		ImageUrl    string   `json:"image_url" validate:"required"`
	}

	SearchBookRequest struct {
		Search string `json:"search" validate:"required"`
	}
)
