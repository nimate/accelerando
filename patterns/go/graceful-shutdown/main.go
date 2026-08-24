package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("worker: shutdown requested")
			return
		case <-ticker.C:
			doWork(ctx)
		}
	}
}

func doWork(ctx context.Context) {
	log.Println("worker: tick")
	_ = ctx
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	worker(ctx)
	log.Println("main: exiting cleanly")
}
