package handlers

import (
	handlers_interfaces "github.com/RandySteven/Library-GO/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/Library-GO/interfaces/usecases"
	"net/http"
)

type EventHandler struct {
	usecase usecases_interfaces.EventUsecase
}

func (e *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
}

func (e *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
}

func (e *EventHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
}

var _ handlers_interfaces.EventHandler = &EventHandler{}

func newEventHandler(usecase usecases_interfaces.EventUsecase) *EventHandler {
	return &EventHandler{
		usecase: usecase,
	}
}
