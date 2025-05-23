package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/superdumb33/qoute_service/internal/app"
	"github.com/superdumb33/qoute_service/internal/config"
)

func main() {
	cfg := config.MustInit()

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app := app.New(cfg, log)

	go app.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<- stop

	log.Info("Gracefully stopping server...")
	app.Stop(context.Background())
	log.Info("Server stopped")
}