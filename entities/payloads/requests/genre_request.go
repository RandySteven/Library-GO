package requests

type (
	GenreRequest struct {
		Genre string `json:"genre" validate:"required"`
	}

	ChooseFavoriteGenresRequest struct {
		GenreIDs []uint64 `json:"genre_ids" validate:"required"`
	}
)
