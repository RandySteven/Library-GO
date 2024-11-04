package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Entering logging middleware")

		requestTime := time.Now()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)

		var requestBody interface{}
		if err := json.Unmarshal(body, &requestBody); err != nil {
			log.Printf("Failed to unmarshal request body: %v\n", err)
			return
		}

		log.Printf("%s %s %s %s\n", requestTime.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, requestBody)
	})
}
