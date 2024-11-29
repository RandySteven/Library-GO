package middlewares

import (
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
)

type WhitelistedMiddleware struct {
	whitelist map[enums.Middleware]map[string]bool
}

func NewWhitelistedMiddleware() *WhitelistedMiddleware {
	return &WhitelistedMiddleware{
		whitelist: make(map[enums.Middleware]map[string]bool),
	}
}

func (w *WhitelistedMiddleware) RegisterMiddleware(prefix enums.RouterPrefix, path string, middlewares []enums.Middleware) {
	if w == nil {
		_ = NewWhitelistedMiddleware()
	}
	for _, middleware := range middlewares {
		if w.whitelist[middleware] == nil {
			w.whitelist[middleware] = make(map[string]bool)
		}
		w.whitelist[middleware][fmt.Sprintf("%s%s", prefix.ToString(), path)] = true
	}
}

func (w *WhitelistedMiddleware) WhiteListed(uri string, middleware enums.Middleware) bool {
	return w.whitelist[middleware][uri]
}
