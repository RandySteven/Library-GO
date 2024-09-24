package handlers_interfaces

import "net/http"

type StoryGeneratorHandler interface {
	GenerateStory(w http.ResponseWriter, r *http.Request)
}
