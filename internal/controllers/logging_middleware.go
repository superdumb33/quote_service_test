package controllers

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		log.Debug("HTTP request result",
			"method", r.Method,
			"url", r.URL.Path,
			"latency", time.Since(start),
		)
	})
}
