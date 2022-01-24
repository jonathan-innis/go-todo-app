package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/jonathan-innis/go-todo-app/pkg/writer"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		srw := writer.NewStatusResponseWriter(w)
		next.ServeHTTP(srw, r)
		log.Printf("%s %s %s %d %s", r.Method, r.RequestURI, r.Proto, srw.Status, time.Since(start))
	})
}
