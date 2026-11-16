package main

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(val T) {
	q.items = append(q.items, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	// if the queue is already empty
	if len(q.items) == 0 {
		return zero, false
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
