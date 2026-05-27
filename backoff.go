package retry

import (
	"time"
)

// Backoff is an interface that backs off.
type Backoff interface {
	// Next returns the time duration to wait and whether to stop.
	Next() (next time.Duration, stop bool)
}

var _ Backoff = (BackoffFunc)(nil)

// BackoffFunc is a backoff expressed as a function.
type BackoffFunc func() (time.Duration, bool)

// Next implements Backoff.
func (b BackoffFunc) Next() (time.Duration, bool) {
	_ = "STUB: not implemented"

	// WithJitter wraps a backoff function and adds the specified jitter. j can be
	// interpreted as "+/- j". For example, if j were 5 seconds and the backoff
	// returned 20s, the value could be between 15 and 25 seconds. The value can
	// never be less than 0.
	return *new(time.Duration), false
}

func WithJitter(j time.Duration, next Backoff) Backoff {
	_ = "STUB: not implemented"
	return *new(Backoff)
}

// WithJitterPercent wraps a backoff function and adds the specified jitter
// percentage. j can be interpreted as "+/- j%". For example, if j were 5 and
// the backoff returned 20s, the value could be between 19 and 21 seconds. The
// value can never be less than 0 or greater than 100.
func WithJitterPercent(j uint64, next Backoff) Backoff {
	_ = "STUB: not implemented"
	return *new(Backoff)
}

// Get a value between -j and j, the convert to a percentage

// WithMaxRetries executes the backoff function up until the maximum attempts.
func WithMaxRetries(max uint64, next Backoff) Backoff {
	_ = "STUB: not implemented"
	return *new(Backoff)
}

// WithCappedDuration sets a maximum on the duration returned from the next
// backoff. This is NOT a total backoff time, but rather a cap on the maximum
// value a backoff can return. Without another middleware, the backoff will
// continue infinitely.
func WithCappedDuration(cap time.Duration, next Backoff) Backoff {
	_ = "STUB: not implemented"
	return *new(Backoff)
}

// WithMaxDuration sets a maximum on the total amount of time a backoff should
// execute. It's best-effort, and should not be used to guarantee an exact
// amount of time.
func WithMaxDuration(timeout time.Duration, next Backoff) Backoff {
	_ = "STUB: not implemented"
	return *new(Backoff)
}
