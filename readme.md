# Queue

![go report card](https://goreportcard.com/badge/github.com/brian-yu/go-queue)

Package `queue` provides a simple and efficient queue that can hold values of any type.

See GoDoc for futher documentation and examples.
- https://godoc.org/github.com/brian-yu/go-queue

## Example Usage

```go
q := queue.New()

q.Push("hello")
q.Push(32)
q.Push(12.32)
q.Push(author{name: "Brian", age: 19})

t, err := q.Top() // t = "hello", err = nil
q.Len()           // 4

p, err := q.Pop() // p = "hello", err = nil
q.Len()           // 3

t, err := q.Top() // t = 32, err = nil

q.Pop()           // Pop 32
q.Pop()           // Pop 12.32

a, err := q.Pop() // a.(author).name = "Brian", a.(author).age = 19, err = nil
q.Len()           // 0

_, err := q.Pop() // err != nil -- cannot Pop() from empty queue

```

## Installation
```bash
go get github.com/brian-yu/go-queue
```

## Benchmarks
```bash
$ go test -bench=. -benchtime=10000000x -benchmem
goos: darwin
goarch: amd64
pkg: github.com/brian-yu/go-queue
BenchmarkPushInt-8          10000000           100 ns/op          82 B/op          0 allocs/op
BenchmarkPopInt-8           10000000             1.56 ns/op        0 B/op          0 allocs/op
BenchmarkTopInt-8           10000000             0.36 ns/op        0 B/op          0 allocs/op
BenchmarkPushStr-8          10000000            81.2 ns/op        82 B/op          0 allocs/op
BenchmarkPopStr-8           10000000             1.64 ns/op        0 B/op          0 allocs/op
BenchmarkTopStr-8           10000000             0.36 ns/op        0 B/op          0 allocs/op
BenchmarkPushStruct-8       10000000           139 ns/op         114 B/op          1 allocs/op
BenchmarkPopStruct-8        10000000             1.52 ns/op        0 B/op          0 allocs/op
BenchmarkTopStruct-8        10000000             0.34 ns/op        0 B/op          0 allocs/op
BenchmarkPushMixed-8        10000000           143 ns/op          93 B/op          0 allocs/op
BenchmarkPopMixed-8         10000000             1.55 ns/op        0 B/op          0 allocs/op
BenchmarkTopMixed-8         10000000             0.39 ns/op        0 B/op          0 allocs/op
```