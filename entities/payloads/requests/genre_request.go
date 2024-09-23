package requests

type (
	GenreRequest struct {
		Genre string `json:"genre" validate:"required"`
	}
)
