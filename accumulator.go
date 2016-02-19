// Package accumulator is a lightweight wrapper around an atomic.Int64, which
// allows you to easily increment a counter, and "flush" it to 0, returning the
// current value. A use case for this might be a timed reading in a stats system.
package accumulator

import "sync/atomic"

// Ac is a wrapper around an *int64, interacted with via the atomic package
type Ac struct {
	n *int64
}

// New returns a new pointer to an accumulator
func New() *Ac {
	i := int64(0)
	return &Ac{n: &i}
}

// Incr incremenets the accumulator by 1
func (a *Ac) Incr() {
	atomic.AddInt64(a.n, 1)
}

// Flush resets the accumulator to 0, and returns the former value
func (a *Ac) Flush() int64 {
	return atomic.SwapInt64(a.n, 0)
}
