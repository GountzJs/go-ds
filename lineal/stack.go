package lineal

type stack struct {
	items    []int
	capacity uint
}

func NewStack(capacity uint) *stack {
	return &stack{
		capacity: capacity,
		items:    make([]int, capacity),
	}
}

func (s stack) IsFull() bool {
	return len(s.items) == int(s.capacity)
}

func (s stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) Push(item int) {
	if s.IsFull() {
		return
	}

	s.items = append(s.items, item)
}

func (s *stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	var lastIndex = len(s.items) - 1

	var lastValue = s.items[lastIndex]

	s.items = s.items[:lastIndex]

	return lastValue
}

func (s stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	var lastIndex = len(s.items) - 1

	return s.items[lastIndex]
}
