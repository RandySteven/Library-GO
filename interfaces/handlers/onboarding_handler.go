package handlers_interfaces

import "net/http"

type OnboardingHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	VerifyUser(w http.ResponseWriter, r *http.Request)
	GetLoginUser(w http.ResponseWriter, r *http.Request)
}
