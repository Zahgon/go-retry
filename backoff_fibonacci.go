package retry

import (
	"context"
	"time"
	"unsafe"
)

type state [2]time.Duration

type fibonacciBackoff struct {
	state unsafe.Pointer
}

// Fibonacci is a wrapper around Retry that uses a Fibonacci backoff. See
// NewFibonacci.
func Fibonacci(ctx context.Context, base time.Duration, f RetryFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// NewFibonacci creates a new Fibonacci backoff using the starting value of
// base. The wait time is the sum of the previous two wait times on each failed
// attempt (1, 1, 2, 3, 5, 8, 13...).
//
// Once it overflows, the function constantly returns the maximum time.Duration
// for a 64-bit integer.
//
// It panics if the given base is less than zero.
func NewFibonacci(base time.Duration) Backoff { _ = "STUB: not implemented"; return *new(Backoff) }

// Next implements Backoff. It is safe for concurrent use.
func (b *fibonacciBackoff) Next() (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}
