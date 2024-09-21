package routes

import (
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/handlers"
	"net/http"
)

type (
	HandlerFunc func(w http.ResponseWriter, r *http.Request)

	Router struct {
		path    string
		handler HandlerFunc
		method  string
	}
)

func RegisterEndpointRouter(path, method string, handler HandlerFunc) *Router {
	return &Router{path: path, handler: handler, method: method}
}

func NewEndpointRouters(h *handlers.Handlers) map[enums.RouterPrefix][]*Router {
	endpointRouters := make(map[enums.RouterPrefix][]*Router)

	endpointRouters[enums.OnboardingPrefix] = []*Router{
		RegisterEndpointRouter("/register", http.MethodPost, h.OnboardingHandler.RegisterUser),
		RegisterEndpointRouter("/login", http.MethodPost, h.OnboardingHandler.LoginUser),
		RegisterEndpointRouter("/verify", http.MethodPost, h.OnboardingHandler.VerifyUser),
	}

	endpointRouters[enums.UserPrefix] = []*Router{}

	endpointRouters[enums.BookPrefix] = []*Router{}

	return endpointRouters
}
