package routes

import "net/http"

func registerEndpointRouter(path, method string, handler HandlerFunc) *Router {
	return &Router{path: path, handler: handler, method: method}
}

func Post(path string, handler HandlerFunc) *Router {
	return registerEndpointRouter(path, http.MethodPost, handler)
}

func Put(path string, handler HandlerFunc) *Router {
	return registerEndpointRouter(path, http.MethodPut, handler)
}

func Delete(path string, handler HandlerFunc) *Router {
	return registerEndpointRouter(path, http.MethodDelete, handler)
}

func Get(path string, handler HandlerFunc) *Router {
	return registerEndpointRouter(path, http.MethodGet, handler)
}
