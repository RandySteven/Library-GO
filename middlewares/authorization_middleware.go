package middlewares

import (
	"github.com/RandySteven/Library-GO/enums"
	"github.com/RandySteven/Library-GO/utils"
	"net/http"
)

type route struct {
	endpoint string
	method   string
}

func (mv *MiddlewareValidator) AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !mv.whitelist.WhiteListed(r.Method, utils.ReplaceLastURLID(r.RequestURI), enums.AuthenticationMiddleware) {
			next.ServeHTTP(w, r)
			return
		}

		whiteListedRouting := make(map[enums.RoleUser][]*route)
		whiteListedRouting[enums.Admin] = []*route{
			{
				endpoint: `/books`,
				method:   http.MethodGet,
			},
			{
				endpoint: `/books`,
				method:   http.MethodPost,
			},
			{
				endpoint: `/genres`,
				method:   http.MethodPost,
			},
			{
				endpoint: `/genres`,
				method:   http.MethodGet,
			},
		}

		//ctx := r.Context()
		//roleID := ctx.Value(enums.RoleID).(uint64)

		next.ServeHTTP(w, r)
	})
}
