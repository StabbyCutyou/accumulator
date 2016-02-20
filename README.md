# Accumulator [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/StabbyCutyou/accumulator)

Accumulator is a collection of threadsafe counters that you can increment, and collect
periodic results from.

This is probably the definition of death by a 1,000 libraries, but I wanted something
to collect statistics with, that were quick, threadsafe, and easy to re-use and share.

Accumulator provides a simple struct that lets you increment a counter, and flush it
to zero, returning the current value.

# Using it

Download the library

```go
go get "github.com/StabbyCutyou/buffstreams"
```

Import the library

```go
import "github.com/StabbyCutyou/buffstreams"
```

There are two types of counters, currently

## Int64

Int64 is a wrapper around the atomic packages AddInt64 and SwapInt64 calls.

```go
a := accumulator.NewInt64()
go func() {
  for {
    a.Incr()
  }
}()
go func() {
    for{
      fmt.Println(a.Flush())
    }
}
```

#Float64

Float64 is a wrapper around a Mutex that controls a float64.

```go
a := accumulator.NewFloat64()
go func() {
  for {
    a.Incr()
  }
}()
go func() {
    for{
      fmt.Println(a.Flush())
    }
}
```

# Roadmap

* Figure out how to use go generate so I can build out one for each for numeric type

LICENSE
=========
Apache v2 - See LICENSE
