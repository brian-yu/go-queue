# Queue
Package `queue` provides a simple and efficient queue that can hold values of any type.

## Example

```go
q := New()

q.Push("hello")
q.Push(32)
q.Push(12.32)
q.Push(author{name: "Brian", age: 19})

t, err := q.Top() // t = "hello", err = nil
q.Len()           // 4

p, err := q.Pop() // p = "hello", err = nil
q.Len()           // 3

t, err := q.Top() // t = 32, err = nil

q.Pop()
q.Pop()

a, err := q.Pop() // a.(author).name = "Brian", a.(author).age = 19, err = nil

_, err := q.Pop() // err != nil -- cannot Pop() from empty queue

```