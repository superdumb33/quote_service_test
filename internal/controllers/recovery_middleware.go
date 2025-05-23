package controllers

import (
	"log/slog"
	"net/http"
)

func RecoveryMiddleware(next http.Handler, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func () {
			err := recover()
			if err != nil {
				log.Error("recovered from panic", "error", err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}