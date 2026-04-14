package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		log.WithFields(log.Fields{
			"method":  r.Method,
			"path":    r.URL.Path,
			"status":  rw.status,
			"elapsed": time.Since(start).String(),
		}).Info("request completed")
	})
}
