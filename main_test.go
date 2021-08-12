package main

import (
	"context"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRateLimiter(t *testing.T) {
	ctx := context.Background()
	numItems := 300000
	batchSize := 1000
	endTime := time.Now().Add(10 * time.Second)
	timeRemaining := time.Until(endTime)
	dispatchInterval := time.Duration(int64(timeRemaining) / int64(numItems/batchSize))
	limiter := rate.NewLimiter(rate.Every(dispatchInterval), batchSize)

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
