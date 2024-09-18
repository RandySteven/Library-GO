package routes

import "net/http"

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

//func NewEndpointRouters(h *Handlers) map[enums.RouterPrefix][]Router {
//	endpointRouters := make(map[enums.RouterPrefix][]Router)
//}
