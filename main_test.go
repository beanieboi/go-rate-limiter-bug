package main

import (
	"context"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func dispatchLimiter(duration time.Duration, numItems int) *rate.Limiter {
	endTime := time.Now().Add(duration)
	timeRemaining := time.Until(endTime)
	dispatchInterval := time.Duration(int64(timeRemaining) / int64(numItems))
	return rate.NewLimiter(rate.Every(dispatchInterval), 1)
}

func TestRateLimiter(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	ctx := context.Background()
	numItems := 300000
	limiter := dispatchLimiter(10*time.Second, numItems)

	loopStart := time.Now()
	for i := 0; i < numItems; i++ {
		err := limiter.Wait(ctx)
		if err != nil {
			t.Errorf("error waiting: %v", err)
		}
		// do actual dispatching
	}
	duration := time.Since(loopStart)

	if duration > 15*time.Second {
		t.Errorf("expected execution time of less than 12, got: %v", duration)
	}
}
