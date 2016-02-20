package accumulator

import (
	"testing"
	"time"
)

func BenchmarkAtomic(b *testing.B) {
	c := NewInt64()
	for i := 0; i < b.N; i++ {
		c.Incr()
	}
}

func TestAtomic(t *testing.T) {
	c := NewInt64()
	for i := 0; i < 100; i++ {
		c.Incr()
	}
	if c.Flush() != 100 {
		t.Fail()
	}
}

func BenchmarkMutex(b *testing.B) {
	c := NewFloat64()
	for i := 0; i < b.N; i++ {
		c.Incr()
	}
}

func TestMutex(t *testing.T) {
	c := NewFloat64()
	for i := 0; i < 100; i++ {
		c.Incr()
	}
	if c.Flush() != 100.0 {
		t.Fail()
	}
}

func BenchmarkAtomicWithFlush(b *testing.B) {
	c := NewInt64()
	for i := 0; i < b.N; i++ {
		c.Incr()
		if i%10000 == 0 {
			c.Flush()
		}
	}
}

func BenchmarkMutexWithFlush(b *testing.B) {
	c := NewFloat64()
	for i := 0; i < b.N; i++ {
		c.Incr()
		if i%10000 == 0 {
			c.Flush()
		}
	}
}

func BenchmarkAtomicWithConcurrentFlush(b *testing.B) {
	c := NewInt64()
	sleep, _ := time.ParseDuration("10ms")
	go func() {
		for {
			time.Sleep(sleep)
			c.Flush()
		}
	}()
	for i := 0; i < b.N; i++ {
		c.Incr()
	}
}

func BenchmarkMutexWithConcurrentFlush(b *testing.B) {
	c := NewFloat64()
	sleep, _ := time.ParseDuration("10ms")
	go func() {
		for {
			time.Sleep(sleep)
			c.Flush()
		}
	}()
	for i := 0; i < b.N; i++ {
		c.Incr()
		if i%10000 == 0 {
			c.Flush()
		}
	}
}
