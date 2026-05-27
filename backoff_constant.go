package retry

import (
	"context"
	"time"
)

// Constant is a wrapper around Retry that uses a constant backoff. It panics if
// the given base is less than zero.
func Constant(ctx context.Context, t time.Duration, f RetryFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// NewConstant creates a new constant backoff using the value t. The wait time
// is the provided constant value. It panics if the given base is less than
// or equal to zero.
func NewConstant(t time.Duration) Backoff { _ = "STUB: not implemented"; return *new(Backoff) }
