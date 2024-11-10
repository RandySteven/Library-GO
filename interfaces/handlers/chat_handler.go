package handlers_interfaces

import "net/http"

type (
	RoomChatHandler interface {
		CreateRoomChat(w http.ResponseWriter, r *http.Request)
		GetAllRoomsChat(w http.Response, r *http.Request)
	}

	SendChatHandler interface {
		SendChat(w http.ResponseWriter, r *http.Request)
	}
)
