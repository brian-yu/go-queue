package queue

import (
	"testing"
)

func checkQueueLen(t *testing.T, q *Queue, expect int) {
	if n := q.Len(); n != expect {
		t.Errorf("q.Len() = %d, want %d", n, expect)
	}
}

func TestNew(t *testing.T) {
	q := New()
	checkQueueLen(t, q, 0)
}

type tester struct {
	name string
	age  int
}

func testQueue() *Queue {
	q := New()

	q.Push("hello")
	q.Push(32)
	q.Push(12.32)
	q.Push(tester{name: "Brian", age: 19})

	return q
}

func TestPush(t *testing.T) {
	q := testQueue()
	checkQueueLen(t, q, 4)
}

func TestPop(t *testing.T) {
	q := testQueue()
	e1 := "hello"
	if v, err := q.Pop(); v != e1 {
		t.Errorf("q.Pop() = %v, %v, want %v, %v", v, err, e1, nil)
	}
	e2 := 32
	if v, err := q.Pop(); v != e2 {
		t.Errorf("q.Pop() = %v, %v, want %v, %v", v, err, e2, nil)
	}
	e3 := 12.32
	if v, err := q.Pop(); v != e3 {
		t.Errorf("q.Pop() = %v, %v, want %v, %v", v, err, e3, nil)
	}
	e4 := tester{name: "Brian", age: 19}
	if v, err := q.Pop(); v.(tester).name != e4.name || v.(tester).age != e4.age {
		t.Errorf("q.Pop() = %v, %v, want %v, %v", v, err, e4, nil)
	}

	if _, err := q.Pop(); err == nil {
		t.Error("q.Pop() on empty queue did not return error")
	}
}

func TestPopEmpty(t *testing.T) {
	q := New()
	if _, err := q.Pop(); err == nil {
		t.Error("q.Pop() on empty queue did not return error")
	}
}

func TestTop(t *testing.T) {
	q := testQueue()

	e1 := "hello"
	if v, err := q.Top(); v != e1 {
		t.Errorf("q.Top() = %v, %v, want %v, %v", v, err, e1, nil)
	}
	if v, err := q.Top(); v != e1 {
		t.Errorf("q.Top() = %v, %v, want %v, %v", v, err, e1, nil)
	}
	q.Pop()
	e2 := 32
	if v, err := q.Top(); v != e2 {
		t.Errorf("q.Top() = %v, %v, want %v, %v", v, err, e2, nil)
	}
	q.Pop()
	q.Pop()
	q.Pop()

	if _, err := q.Top(); err == nil {
		t.Error("q.Top() on empty queue did not return error")
	}
}

func TestTopEmpty(t *testing.T) {
	q := New()
	if _, err := q.Top(); err == nil {
		t.Error("q.Top() on empty queue did not return error")
	}
}
