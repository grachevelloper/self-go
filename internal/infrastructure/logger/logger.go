package logger

import (
	"io"
	"log/slog"
	"strings"
)

type Config struct {
	Env string
}

func New(cfg Config, output io.Writer) *slog.Logger {
	options := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if strings.EqualFold(cfg.Env, "development") {
		return slog.New(slog.NewTextHandler(output, options))
	}

	return slog.New(slog.NewJSONHandler(output, options))
}
