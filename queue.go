// Package queue is a simple yet efficient queue implementation
package queue

import (
	"errors"
	"fmt"
)

// Queue represents the queue and has a variety of methods
// that allow values to be enqueued, dequeued, and examined.
type Queue struct {
	s []interface{}
}

// New creates and returns a Queue
func New() *Queue {
	return &Queue{s: make([]interface{}, 0)}
}

// Push enqueues the supplied value to the end of the queue
func (q *Queue) Push(v interface{}) {
	q.s = append(q.s, v)
}

// Pop dequeues and returns a value from the front of the
// queue or returns an error if the queue is empty
func (q *Queue) Pop() (interface{}, error) {
	if len(q.s) == 0 {
		return nil, errors.New("Queue: cannot Pop() from empty queue")
	}
	res := q.s[0]
	// q.s[0] = struct{}{}
	q.s = q.s[1:]
	return res, nil
}

// Top returns the value at the front of the
// queue or returns an error if the queue is empty
func (q *Queue) Top() (interface{}, error) {
	if len(q.s) == 0 {
		return nil, errors.New("Queue: cannot Top() on empty queue")
	}
	return q.s[0], nil
}

// Len returns the number of elements in the queue
func (q *Queue) Len() interface{} {
	return len(q.s)
}

// String provides a string representation of the queue
func (q Queue) String() string {
	return fmt.Sprintf("%v", q.s)
}
