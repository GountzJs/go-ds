package lineal

import "errors"

type stack[T any] struct {
	items    []T
	capacity uint
}

func NewStack[T any](capacity uint) *stack[T] {
	return &stack[T]{
		capacity: capacity,
		items:    make([]T, 0, capacity),
	}
}

func (s stack[T]) IsFull() bool {
	return len(s.items) == int(s.capacity)
}

func (s stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack[T]) Push(item T) error {
	if s.IsFull() {
		return errors.New("Stack is full")
	}

	s.items = append(s.items, item)
	return nil
}

func (s *stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	var lastIndex = len(s.items) - 1

	var lastValue = s.items[lastIndex]

	s.items = s.items[:lastIndex]

	return lastValue, nil
}

func (s stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	var lastIndex = len(s.items) - 1

	return s.items[lastIndex], nil
}
