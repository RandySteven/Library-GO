package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutTime, err := strconv.Atoi(os.Getenv("SERVER_TIMEOUT"))
		if err != nil {
			log.Printf("Could not parse timeout value: %v", err)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), time.Duration(timeoutTime)*time.Second)
		defer cancel()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		if err := ctx.Err(); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Println(err)
				return
			}
		}
	})
}
