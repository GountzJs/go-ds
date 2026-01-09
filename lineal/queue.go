package lineal

import "errors"

type queue[T any] struct {
	items    []T
	capacity uint
	head     int
	tail     int
}

func NewQueue[T any](capacity uint) *queue[T] {
	return &queue[T]{
		capacity: capacity,
		items:    make([]T, capacity),
		head:     -1,
		tail:     -1,
	}
}

func (q queue[T]) IsEmpty() bool {
	return q.head == -1 && q.tail == -1
}

func (q queue[T]) IsFull() bool {
	return q.tail == int(q.capacity-1)
}

func (q queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("Queue is empty")
	}

	return q.items[q.head], nil
}

func (q queue[T]) Size() int {
	if q.IsEmpty() {
		return 0
	}
	return q.tail - q.head + 1
}

func (q *queue[T]) Enqueue(value T) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}

	if q.head == -1 {
		var initialIndex = 0
		q.head = initialIndex
		q.tail = initialIndex
		q.items[initialIndex] = value
		return nil
	}

	q.tail = q.tail + 1
	q.items[q.tail] = value
	return nil
}

func (q *queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("Queue is empty")
	}

	if q.head == q.tail {
		var resetIndex = -1
		var valueOfIndex = q.items[q.head]
		q.head = resetIndex
		q.tail = resetIndex
		return valueOfIndex, nil
	}

	var valueOfIndex = q.items[q.head]
	q.head = q.head + 1

	return valueOfIndex, nil
}
