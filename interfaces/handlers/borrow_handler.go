package handlers_interfaces

import "net/http"

type BorrowHandler interface {
	BorrowCheckout(w http.ResponseWriter, r *http.Request)
}
