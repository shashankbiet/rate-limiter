package ratelimiter

import (
	"sync"
	"time"
)

type tokenBucket struct {
	rate       time.Duration // Rate of tokens to be added per duration
	capacity   int           // Maximum number of tokens the bucket can hold
	tokens     int           // Current number of tokens in the bucket
	lastUpdate time.Time
	mutex      sync.Mutex // Mutex for thread safety
}

// NewTokenBucket creates a new tokenBucket with the given rate and capacity.
func NewTokenBucket(rate time.Duration, capacity int) *tokenBucket {
	return &tokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastUpdate: time.Now(),
	}
}

// Allow checks if a request is allowed based on the rate limiter.
func (tb *tokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	elapsed := time.Since(tb.lastUpdate)
	tb.lastUpdate = time.Now()
	// tb.tokens = min(tb.tokens+tb.rate*elapsed, tb.capacity)
	// Calculate the number of tokens to add based on the elapsed time
	tokensToAdd := int(elapsed / tb.rate)
	tb.tokens = min(tb.tokens+tokensToAdd, tb.capacity)

	if tb.tokens > 0 {
		tb.tokens -= 1
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
