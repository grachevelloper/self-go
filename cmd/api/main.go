package main

import (
	"book-service/internal/app"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	bootstrapLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	application, err := app.New(ctx)
	if err != nil {
		bootstrapLogger.Error("init app failed", "error", err)
		os.Exit(1)
	}
	defer application.Close()

	go func() {
		application.Logger.Info("http server started", "addr", application.Server.Addr)

		if err := application.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			application.Logger.Error("http server failed", "error", err)
			os.Exit(1)
		}
	}()

	waitForShutdown(application)

}
