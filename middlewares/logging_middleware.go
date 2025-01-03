package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/RandySteven/Library-GO/enums"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriterWrapper) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			body []byte = nil
			err  error  = nil
			ctx         = context.WithValue(r.Context(), enums.RequestID, uuid.NewString())
		)
		requestTime := time.Now()

		if r.Method != http.MethodGet {
			body, err = ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("Failed to read request body: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		rw := &ResponseWriterWrapper{
			ResponseWriter: w,
			Body:           &bytes.Buffer{},
			StatusCode:     http.StatusOK,
		}

		r.WithContext(ctx)
		next.ServeHTTP(rw, r)

		var requestBody interface{}
		if body != nil {
			if err := json.Unmarshal(body, &requestBody); err != nil {
				log.Printf("Failed to unmarshal request body: %v\n", err)
				return
			}
		}

		var responseBody interface{}
		if rw.Body.Len() > 0 {
			if err := json.Unmarshal(rw.Body.Bytes(), &responseBody); err != nil {
				log.Printf("Failed to unmarshal response body: %v\n", err)
			}
		}

		logEntry := map[string]interface{}{
			"Request ID ": ctx.Value(enums.RequestID).(string),
			"timestamp":   requestTime.Format("2006-01-02 15:04:05"),
			"method":      r.Method,
			"path":        r.URL.Path,
			"remote":      r.RemoteAddr,
			"request":     requestBody,
			"response":    responseBody,
			"status":      rw.StatusCode,
		}

		logEntryJSON, err := json.Marshal(logEntry)
		if err != nil {
			log.Printf("Failed to marshal log entry: %v\n", err)
			return
		}

		log.Println(string(logEntryJSON))
	})
}

//func writeLog() {
//	day, month, year := time.Now().Date()
//	dateFile := fmt.Sprintf("%d%d%d.log", year, month, day)
//	var logFileCh = make(chan *os.File)
//	_, err := os.Open(dateFile)
//	if err != nil {
//		if errors.As(err, fs.ErrNotExist) {
//			defer func() {
//				logFile, err := os.Create(dateFile)
//				if err != nil {
//					return
//				}
//				logFileCh <- logFile
//			}()
//		}
//	}
//
//	logFile = <-logFileCh
//}
