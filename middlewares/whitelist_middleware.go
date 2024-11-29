package middlewares

import (
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
)

type (
	WhitelistedMiddleware struct {
		whitelist map[enums.Middleware]map[string]bool
	}

	MiddlewareValidator struct {
		whitelist *WhitelistedMiddleware
	}
)

func NewMiddlewareValidator(whitelist *WhitelistedMiddleware) *MiddlewareValidator {
	return &MiddlewareValidator{whitelist: whitelist}
}

func NewWhitelistedMiddleware() *WhitelistedMiddleware {
	return &WhitelistedMiddleware{
		whitelist: make(map[enums.Middleware]map[string]bool),
	}
}

func (w *WhitelistedMiddleware) RegisterMiddleware(prefix enums.RouterPrefix, method string, path string, middlewares []enums.Middleware) {
	if w == nil {
		_ = NewWhitelistedMiddleware()
	}
	if middlewares == nil {
		return
	}
	for _, middleware := range middlewares {
		if w.whitelist[middleware] == nil {
			w.whitelist[middleware] = make(map[string]bool)
		}
		w.whitelist[middleware][fmt.Sprintf("%s|%s%s", method, prefix.ToString(), path)] = true
	}
}

func (w *WhitelistedMiddleware) WhiteListed(method string, uri string, middleware enums.Middleware) bool {
	return w.whitelist[middleware][fmt.Sprintf("%s|%s", method, uri)]
}
