package accesslog

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func AccessLogMiddleware(
	logger *logrus.Logger,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Request received")

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		fields := logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   rw.status,
			"response": http.StatusText(rw.status),
			"time":     time.Since(start),
		}
		logger.WithFields(fields).Info("Request processed")
	})
}
