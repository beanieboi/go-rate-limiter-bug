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
	burstBuffer := 2
	endTime := time.Now().Add(10 * time.Second)
	timeRemaining := time.Until(endTime)
	dispatchInterval := time.Duration(int64(timeRemaining) / int64(numItems))
	burst := int(time.Millisecond / dispatchInterval)
	limiter := rate.NewLimiter(rate.Every(dispatchInterval), burst*burstBuffer)

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
		t.Errorf("expected execution time of less than 15, got: %v", duration)
	}
}
