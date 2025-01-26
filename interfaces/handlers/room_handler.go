package handlers_interfaces

import "net/http"

type RoomHandler interface {
	AddNewRoom(w http.ResponseWriter, r *http.Request)
	GetRooms(w http.ResponseWriter, r *http.Request)
	GetRoomByID(w http.ResponseWriter, r *http.Request)
	UploadRoom(w http.ResponseWriter, r *http.Request)
}
