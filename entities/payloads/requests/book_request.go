package requests

type (
	CreateBookRequest struct {
		Title       string   `form:"title" validate:"required"`
		Description string   `form:"description" validate:"required"`
		Genres      []uint64 `form:"genres" validate:"required"`
		Authors     []uint64 `form:"authors" validate:"required"`
		Image       string   `form:"image_url" validate:"required"`
	}

	SearchBookRequest struct {
		Search string `json:"search" validate:"required"`
	}
)
