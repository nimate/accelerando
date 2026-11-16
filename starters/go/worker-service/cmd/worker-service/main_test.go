package main

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"testing/synctest"
	"time"
)

func TestPollerStopsWhileWaitingForToken(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		tokens := make(chan time.Time)
		// Release a stuck receiver even if the cancellation assertion fails.
		defer close(tokens)

		logger := slog.New(slog.NewTextHandler(io.Discard, nil))
		result := make(chan error, 1)
		go func() {
			result <- poller(ctx, logger, time.Millisecond, tokens)
		}()

		// Wait for ticker creation, advance fake time, then wait for the token receive.
		synctest.Wait()
		time.Sleep(time.Millisecond)
		synctest.Wait()
		cancel()
		synctest.Wait()

		select {
		case err := <-result:
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("poller returned %v, want context cancellation", err)
			}
		default:
			t.Fatal("poller did not stop while waiting for a token")
		}
	})
}

func TestBurstLimiterStopsAfterCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	_, done := newBurstLimiter(ctx, 1, time.Hour)
	cancel()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("limiter refill loop did not stop after cancellation")
	}
}
