package main

import (
	"context"
	"testing"
	"time"
)

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
