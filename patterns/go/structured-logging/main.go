package main

import (
	"log/slog"
	"os"
)

func newLogger(service string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service", service)
}

func main() {
	logger := newLogger("demo-service")
	logger.Info("service started", "port", 8080)
}
