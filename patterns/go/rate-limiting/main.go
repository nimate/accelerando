package main

import (
	"context"
	"fmt"
	"time"
)

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
	ctx, cancel := context.WithCancel(context.Background())
	limiter, done := newBurstLimiter(ctx, 3, 200*time.Millisecond)

	start := time.Now()
	for requestID := range 8 {
		<-limiter
		fmt.Printf("request %d served after %s\n", requestID+1, time.Since(start).Round(time.Millisecond))
	}
	cancel()
	<-done
}
