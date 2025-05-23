package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/superdumb33/qoute_service/internal/config"
	"github.com/superdumb33/qoute_service/internal/controllers"
	"github.com/superdumb33/qoute_service/internal/infrastructure/database"
	"github.com/superdumb33/qoute_service/internal/infrastructure/repository/pgxrepo"
	"github.com/superdumb33/qoute_service/internal/services"
)

type App struct {
	server *http.Server
	port   int
	log    *slog.Logger
}

func New(cfg config.AppCfg, log *slog.Logger) *App {
	pool := database.MustInitNewPool(cfg)
	repo := pgxrepo.NewPgxQuoteRepo(pool, log)
	service := services.NewQuoteService(repo, log)
	handler := controllers.NewQuoteController(service, log)

	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return controllers.RecoveryMiddleware(next, log)
	})
	router.Use(func(next http.Handler) http.Handler {
		return controllers.LoggingMiddleware(next, log)
	})
	
	handler.RegisterRoutes(router)

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	server := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}

	return &App{server: server, port: cfg.AppPort, log: log}
}

// it'll throw a panic if something goes wrong
func (a *App) MustRun() {
	a.log.Info("Starting server", "app port", a.port)
	if err := a.server.ListenAndServe(); err != nil {
		a.log.Error("Server has failed to start", "error", err)
		panic(err)
	}
}

func (a *App) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
