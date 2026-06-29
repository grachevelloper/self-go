package main

import (
	"book-service/internal/app"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 10 * time.Second

func waitForShutdown(application *app.App) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stop)

	sig := <-stop
	application.Logger.Info("shutdown signal received", "signal", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := application.Server.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		application.Logger.Error("http server shutdown failed", "error", err)
		return
	}

	application.Logger.Info("http server stopped")
}
