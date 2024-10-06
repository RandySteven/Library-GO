package middlewares

import (
	"context"
	"net/http"
	"os"
	"time"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serverTimeout, _ := time.ParseDuration(os.Getenv("SERVER_TIMEOUT"))
		ctx, cancel := context.WithTimeout(r.Context(), serverTimeout*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
