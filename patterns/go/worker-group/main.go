package main

import (
	"context"
	"errors"
	"log"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func tickerWorker(ctx context.Context) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("ticker-worker: stopping")
			return ctx.Err()
		case <-ticker.C:
			log.Println("ticker-worker: tick")
		}
	}
}

func queueWorker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("queue-worker: stopping")
			return ctx.Err()
		case <-time.After(3 * time.Second):
			log.Println("queue-worker: handled message")
		}
	}
}

func main() {
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(rootCtx)
	group.Go(func() error { return tickerWorker(ctx) })
	group.Go(func() error { return queueWorker(ctx) })

	if err := group.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("main: stopping on error: %v", err)
		return
	}

	log.Println("main: workers stopped cleanly")
}
