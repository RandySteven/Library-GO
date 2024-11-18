package middlewares

import (
	"context"
	"github.com/RandySteven/Library-GO/enums"
	caches_client "github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/RandySteven/Library-GO/utils"
	ip "github.com/vikram1565/request-ip"
	"net/http"
)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIp := ip.GetClientIP(r)
		ctx := context.WithValue(r.Context(), enums.ClientIP, clientIp)
		if err := caches_client.RateLimiter(ctx); err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			utils.ResponseHandler(w, http.StatusTooManyRequests, `too many request`, nil, nil, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
