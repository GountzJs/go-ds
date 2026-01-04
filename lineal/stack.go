package lineal

type Stack struct {
	data     []int
	capacity uint
}

func (s *Stack) SetCapacity(value uint) {
	s.capacity = value
}

func (s Stack) IsFull() bool {
	return len(s.data) == int(s.capacity)
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(item int) {
	if s.IsFull() {
		return
	}

	s.data = append(s.data, item)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	var lastIndex = len(s.data) - 1

	var lastValue = s.data[lastIndex]

	s.data = s.data[:lastIndex]

	return lastValue
}

func (s Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	var lastIndex = len(s.data) - 1

	return s.data[lastIndex]
}
