// Package retry provides helpers for retrying.
//
// This package defines flexible interfaces for retrying Go functions that may
// be flakey or eventually consistent. It abstracts the "backoff" (how long to
// wait between tries) and "retry" (execute the function again) mechanisms for
// maximum flexibility. Furthermore, everything is an interface, so you can
// define your own implementations.
//
// The package is modeled after Go's built-in HTTP package, making it easy to
// customize the built-in backoff with your own custom logic. Additionally,
// callers specify which errors are retryable by wrapping them. This is helpful
// with complex operations where only certain results should retry.
package retry

import (
	"context"
)

// RetryFunc is a function passed to [Do].
type RetryFunc func(ctx context.Context) error

// RetryFuncValue is a function passed to [Do] which returns a value.
type RetryFuncValue[T any] func(ctx context.Context) (T, error)

type retryableError struct {
	err error
}

// RetryableError marks an error as retryable.
func RetryableError(err error) error { _ = "STUB: not implemented"; return nil }

// Unwrap implements error wrapping.
func (e *retryableError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error returns the error string.
	return nil
}

func (e *retryableError) Error() string { _ = "STUB: not implemented"; return "" }

func DoValue[T any](ctx context.Context, b Backoff, f RetryFuncValue[T]) (T, error) {
	_ = "STUB: not implemented"

	// Return immediately if ctx is canceled
	return *new(T), nil
}

// Not retryable

// Wait until next attempt or until the context expires. Any error will
// be caught at the top of the loop.

// Do wraps a function with a backoff to retry. The provided context is the same
// context passed to the [RetryFunc].
func Do(ctx context.Context, b Backoff, f RetryFunc) error { _ = "STUB: not implemented"; return nil }
