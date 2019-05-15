package queue

import (
	"math/rand"
	"testing"
)

func exampleQueueInt(n int) *Queue {
	q := New()
	for i := 0; i < n; i++ {
		q.Push(42)
	}
	return q
}

func BenchmarkPushInt(b *testing.B) {
	exampleQueueInt(b.N)
}

func BenchmarkPopInt(b *testing.B) {
	q := exampleQueueInt(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkTopInt(b *testing.B) {
	q := exampleQueueInt(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Top()
	}
}

func exampleQueueStr(n int) *Queue {
	q := New()
	for i := 0; i < n; i++ {
		q.Push("hullo")
	}
	return q
}

func BenchmarkPushStr(b *testing.B) {
	exampleQueueStr(b.N)
}

func BenchmarkPopStr(b *testing.B) {
	q := exampleQueueStr(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkTopStr(b *testing.B) {
	q := exampleQueueStr(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Top()
	}
}

func exampleQueueStruct(n int) *Queue {
	q := New()
	for i := 0; i < n; i++ {
		q.Push(tester{name: "Brian", age: 19})
	}
	return q
}

func BenchmarkPushStruct(b *testing.B) {
	exampleQueueStruct(b.N)
}

func BenchmarkPopStruct(b *testing.B) {
	q := exampleQueueStruct(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkTopStruct(b *testing.B) {
	q := exampleQueueStruct(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Top()
	}
}

func exampleQueueMixed(n int) *Queue {
	q := New()
	for i := 0; i < n; i++ {
		switch rand.Intn(3) {
		case 0:
			q.Push(42)
		case 1:
			q.Push("heyo")
		case 2:
			q.Push(tester{name: "Brian", age: 19})
		}
	}
	return q
}

func BenchmarkPushMixed(b *testing.B) {
	exampleQueueMixed(b.N)
}

func BenchmarkPopMixed(b *testing.B) {
	q := exampleQueueMixed(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkTopMixed(b *testing.B) {
	q := exampleQueueMixed(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Top()
	}
}
