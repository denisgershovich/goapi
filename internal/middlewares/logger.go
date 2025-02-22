package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		loggingWriter := &LoggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		fmt.Printf("Incoming Request: %s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(loggingWriter, r)

		duration := time.Since(start)
		fmt.Printf("Completed Request: %s %s - Status: %d - Duration: %v\n", r.Method, r.URL.Path, loggingWriter.statusCode, duration)
	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}
