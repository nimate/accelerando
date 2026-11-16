package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Config struct {
	ServiceName string
	PollEvery   time.Duration
}

func loadConfig() (Config, error) {
	cfg := Config{
		ServiceName: "worker-service",
		PollEvery:   2 * time.Second,
	}

	if raw := os.Getenv("SERVICE_NAME"); raw != "" {
		cfg.ServiceName = raw
	}

	if raw := os.Getenv("POLL_SECS"); raw != "" {
		secs, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("parse POLL_SECS: %w", err)
		}
		cfg.PollEvery = time.Duration(secs) * time.Second
	}

	return cfg, nil
}

func newLogger(service string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service", service)
}

func newBurstLimiter(ctx context.Context, depth int, rate time.Duration) (<-chan time.Time, <-chan struct{}) {
	limiter := make(chan time.Time, depth)
	done := make(chan struct{})
	for i := 0; i < depth; i++ {
		limiter <- time.Now()
	}

	go func() {
		defer close(done)
		ticker := time.NewTicker(rate)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case tick := <-ticker.C:
				select {
				case limiter <- tick:
				default:
				}
			}
		}
	}()

	return limiter, done
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	logger := newLogger(cfg.ServiceName)
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(rootCtx)
	limiter, limiterDone := newBurstLimiter(ctx, 3, 500*time.Millisecond)

	group.Go(func() error { return poller(ctx, logger, cfg.PollEvery, limiter) })
	group.Go(func() error { return queueWorker(ctx, logger) })

	err = group.Wait()
	<-limiterDone
	if err != nil && !errors.Is(err, context.Canceled) {
		logger.Error("worker group failed", "error", err)
		os.Exit(1)
	}

	logger.Info("service stopped cleanly")
}

func poller(ctx context.Context, logger *slog.Logger, every time.Duration, limiter <-chan time.Time) error {
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("poller stopping")
			return ctx.Err()
		case <-ticker.C:
			select {
			case <-ctx.Done():
				logger.Info("poller stopping")
				return ctx.Err()
			case <-limiter:
			}
			logger.Info("poller tick")
		}
	}
}

func queueWorker(ctx context.Context, logger *slog.Logger) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("queue worker stopping")
			return ctx.Err()
		case <-time.After(3 * time.Second):
			logger.Info("queue worker handled job")
		}
	}
}
