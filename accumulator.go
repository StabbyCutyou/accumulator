// Package accumulator is a lightweight wrapper around an atomic.Int64, which
// allows you to easily increment a counter, and "flush" it to 0, returning the
// current value. A use case for this might be a timed reading in a stats system.
package accumulator

import (
	"sync"
	"sync/atomic"
)

// Int64 is a wrapper around an *int64, interacted with via the atomic package
type Int64 struct {
	n *int64
}

// NewInt64 returns a new pointer to an accumulator
func NewInt64() *Int64 {
	i := int64(0)
	return &Int64{n: &i}
}

// Incr incremenets the accumulator by 1
func (i *Int64) Incr() {
	atomic.AddInt64(i.n, 1)
}

// IncrN incremenets the accumulator by the provided value n
func (i *Int64) IncrN(n int64) {
	atomic.AddInt64(i.n, n)
}

// Flush resets the accumulator to 0, and returns the former value
func (i *Int64) Flush() int64 {
	return atomic.SwapInt64(i.n, 0)
}

// Float64 is a wrapper around an *int64, interacted with via the atomic package
type Float64 struct {
	l sync.Mutex
	n float64
}

// NewFloat64 returns a new pointer to an accumulator
func NewFloat64() *Float64 {
	return &Float64{n: 0}
}

// Incr incremenets the accumulator by 1
func (i *Float64) Incr() {
	i.l.Lock()
	i.n++
	i.l.Unlock()
}

// IncrN incremenets the accumulator by 1
func (i *Float64) IncrN(n float64) {
	i.l.Lock()
	i.n += n
	i.l.Unlock()
}

// Flush resets the accumulator to 0, and returns the former value
func (i *Float64) Flush() float64 {
	i.l.Lock()
	n := i.n
	i.n = 0
	i.l.Unlock()
	return n
}
