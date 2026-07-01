package app

import (
	"book-service/internal/app/server"
	"book-service/internal/config"
	"book-service/internal/delivery"
	"book-service/internal/delivery/http/book"
	"book-service/internal/infrastructure/logger"
	"book-service/internal/infrastructure/postgres"
	postgresbook "book-service/internal/infrastructure/postgres/book"
	"book-service/internal/infrastructure/uuid"
	usecasebook "book-service/internal/usecase/book"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	Config *config.Config
	DB     *sql.DB
	Logger *slog.Logger
	Server *http.Server
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	log := logger.New(logger.Config{Env: cfg.App.Env}, os.Stdout)

	db, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("init postgres: %w", err)
	}

	bookRepository := postgresbook.NewRepository(db)
	bookService := usecasebook.NewUseCase(bookRepository, uuid.GenerateUUID)
	bookHandler := book.NewHandler(bookService, log)

	router := delivery.NewRouter(
		cfg.HTTP.Origin,
		bookHandler,
	)

	httpServer := server.NewHTTPServer(
		cfg.HTTP.Port,
		router,
	)

	return &App{
		Config: cfg,
		DB:     db,
		Logger: log,
		Server: httpServer,
	}, nil
}

func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}

	return nil
}
