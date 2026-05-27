package retry

import (
	"math/rand"
	"sync"
)

type lockedSource struct {
	src *rand.Rand
	mu  sync.Mutex
}

var _ rand.Source64 = (*lockedSource)(nil)

func newLockedRandom(seed int64) *lockedSource { _ = "STUB: not implemented"; return nil }

// Int63 mimics math/rand.(*Rand).Int63 with mutex locked.
func (r *lockedSource) Int63() int64 { _ = "STUB: not implemented"; return 0 }

// Seed mimics math/rand.(*Rand).Seed with mutex locked.
func (r *lockedSource) Seed(seed int64) { _ = "STUB: not implemented"; return }

// Uint64 mimics math/rand.(*Rand).Uint64 with mutex locked.
func (r *lockedSource) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// Int63n mimics math/rand.(*Rand).Int63n with mutex locked.
func (r *lockedSource) Int63n(n int64) int64 { _ = "STUB: not implemented"; return 0 }

// n is power of two, can mask
