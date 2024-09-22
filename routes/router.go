package routes

import (
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/handlers"
	"github.com/RandySteven/Library-GO/middlewares"
	"github.com/gorilla/mux"
	"log"
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

	endpointRouters[enums.BookPrefix] = []*Router{
		RegisterEndpointRouter("", http.MethodPost, h.BookHandler.AddBook),
		RegisterEndpointRouter("", http.MethodGet, h.BookHandler.GetAllBooks),
		RegisterEndpointRouter("/{id}", http.MethodGet, h.BookHandler.GetBookByID),
	}

	return endpointRouters
}

func InitRouters(routers map[enums.RouterPrefix][]*Router, r *mux.Router) {

	onboardingRouter := r.PathPrefix(enums.OnboardingPrefix.ToString()).Subrouter()
	for _, router := range routers[enums.OnboardingPrefix] {
		router.RouterLog(enums.OnboardingPrefix.ToString())
		onboardingRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}

	bookRouter := r.PathPrefix(enums.BookPrefix.ToString()).Subrouter()
	bookRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range routers[enums.BookPrefix] {
		router.RouterLog(enums.BookPrefix.ToString())
		bookRouter.HandleFunc(router.path, router.handler).Methods(router.method)
	}
}

func (router *Router) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
