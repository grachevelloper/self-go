package logger

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestNewUsesTextHandlerForDevelopment(t *testing.T) {
	var out bytes.Buffer

	log := New(Config{Env: "development"}, &out)
	log.Info("service started", slog.String("component", "api"))

	output := out.String()
	if !strings.Contains(output, "msg=\"service started\"") {
		t.Fatalf("expected text log output, got %q", output)
	}
}

func TestNewUsesJSONHandlerOutsideDevelopment(t *testing.T) {
	var out bytes.Buffer

	log := New(Config{Env: "production"}, &out)
	log.Info("service started", slog.String("component", "api"))

	output := out.String()
	if !strings.Contains(output, `"msg":"service started"`) {
		t.Fatalf("expected json log output, got %q", output)
	}
}
