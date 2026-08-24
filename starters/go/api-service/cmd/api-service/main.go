package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Config struct {
	Addr         string
	ServiceName  string
	ShutdownSecs int
}

func loadConfig() (Config, error) {
	cfg := Config{
		Addr:         ":8080",
		ServiceName:  "api-service",
		ShutdownSecs: 10,
	}

	if raw := os.Getenv("PORT"); raw != "" {
		port, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("parse PORT: %w", err)
		}
		cfg.Addr = ":" + strconv.Itoa(port)
	}

	if raw := os.Getenv("SERVICE_NAME"); raw != "" {
		cfg.ServiceName = raw
	}

	if raw := os.Getenv("SHUTDOWN_SECS"); raw != "" {
		secs, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("parse SHUTDOWN_SECS: %w", err)
		}
		cfg.ShutdownSecs = secs
	}

	return cfg, nil
}

func newLogger(service string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service", service)
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	logger := newLogger(cfg.ServiceName)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           requestLoggingMiddleware(logger, mux),
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("http server starting", "addr", cfg.Addr)
		if serveErr := server.ListenAndServe(); serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
			serverErr <- serveErr
		}
		close(serverErr)
	}()

	select {
	case <-ctx.Done():
		logger.Info("shutdown requested")
	case serveErr := <-serverErr:
		if serveErr != nil {
			logger.Error("http server failed", "error", serveErr)
			os.Exit(1)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownSecs)*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("http server shutdown failed", "error", err)
		os.Exit(1)
	}

	logger.Info("service stopped cleanly")
}

func requestLoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
