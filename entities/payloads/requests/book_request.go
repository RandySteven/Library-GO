package requests

import (
	"io"
)

type (
	CreateBookRequest struct {
		Title       string    `form:"title" json:"title" validate:"required"`
		Description string    `form:"description" json:"description" validate:"required"`
		Genres      []uint64  `form:"genres" json:"genres" validate:"required"`
		Authors     []uint64  `form:"authors" json:"authors" validate:"required"`
		Image       io.Reader `form:"image" json:"image" validate:"required"` // Adjusted type for file upload
	}

	SearchBookRequest struct {
		Search string `json:"search" validate:"required"`
	}
)
