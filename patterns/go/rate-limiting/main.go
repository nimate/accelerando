package main

import (
	"fmt"
	"time"
)

func newBurstLimiter(depth int, rate time.Duration) <-chan time.Time {
	limiter := make(chan time.Time, depth)

	for i := 0; i < depth; i++ {
		limiter <- time.Now()
	}

	go func() {
		ticker := time.NewTicker(rate)
		defer ticker.Stop()

		for tick := range ticker.C {
			select {
			case limiter <- tick:
			default:
			}
		}
	}()

	return limiter
}

func main() {
	limiter := newBurstLimiter(3, 200*time.Millisecond)

	start := time.Now()
	for requestID := range 8 {
		<-limiter
		fmt.Printf("request %d served after %s\n", requestID+1, time.Since(start).Round(time.Millisecond))
	}
}
