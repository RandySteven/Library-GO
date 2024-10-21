package handlers_interfaces

import "net/http"

type RatingHandler interface {
	SubmitRating(w http.ResponseWriter, r *http.Request)
}
