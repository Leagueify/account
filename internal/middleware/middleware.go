package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedResWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqStart := time.Now()

		wrappedResponse := &wrappedResWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrappedResponse, r)

		log.Printf(
			"-- %d %s %s%s (%d ms)",
			wrappedResponse.statusCode, r.Method, r.Host,
			r.RequestURI, time.Since(reqStart).Milliseconds(),
		)
	})
}
