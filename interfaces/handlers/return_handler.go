package handlers_interfaces

import "net/http"

type ReturnHandler interface {
	ReturnBook(w http.ResponseWriter, r *http.Request)
}
