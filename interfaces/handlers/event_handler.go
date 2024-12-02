package handlers_interfaces

import "net/http"

type EventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetAllEvents(w http.ResponseWriter, r *http.Request)
	GetEvent(w http.ResponseWriter, r *http.Request)
}
