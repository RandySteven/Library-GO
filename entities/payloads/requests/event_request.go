package requests

import "io"

type (
	CreateEventRequest struct {
		Title             string    `form:"title"`
		Thumbnail         io.Reader `form:"thumbnail"`
		Price             *uint64   `form:"price"`
		Description       string    `form:"description"`
		ParticipantNumber uint64    `form:"participant_number"`
		Date              string    `form:"date"`
		StartTime         string    `form:"start_time"`
		EndTime           string    `form:"end_time"`
	}
)
