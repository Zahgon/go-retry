package retry

import (
	"context"
	"time"
)

type exponentialBackoff struct {
	base    time.Duration
	attempt uint64
}

// Exponential is a wrapper around Retry that uses an exponential backoff. See
// NewExponential.
func Exponential(ctx context.Context, base time.Duration, f RetryFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// NewExponential creates a new exponential backoff using the starting value of
// base and doubling on each failure (1, 2, 4, 8, 16, 32, 64...), up to max.
//
// Once it overflows, the function constantly returns the maximum time.Duration
// for a 64-bit integer.
//
// It panics if the given base is less than zero.
func NewExponential(base time.Duration) Backoff { _ = "STUB: not implemented"; return *new(Backoff) }

// Next implements Backoff. It is safe for concurrent use.
func (b *exponentialBackoff) Next() (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}
