package handlers_interfaces

import "net/http"

type CommentHandler interface {
	Comment(w http.ResponseWriter, r *http.Request)
	Reply(w http.ResponseWriter, r *http.Request)
	GetBookComment(w http.ResponseWriter, r *http.Request)
}
