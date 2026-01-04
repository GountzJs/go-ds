package lineal

type Queue struct {
	items    []int
	capacity uint
}

func (q *Queue) SetCapacity(value uint) {
	q.capacity = value
}

func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q Queue) IsFull() bool {
	return uint(len(q.items)) == q.capacity
}

func (q Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	var lastIndex = len(q.items) - 1
	var item = q.items[lastIndex]
	return item
}

func (q Queue) Enqueue(value int) {
	if q.IsEmpty() {
		return
	}

	var temp = []int{value}

	q.items = append(temp, q.items...)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	var lastIndex = len(q.items) - 1

	var item = q.items[lastIndex]

	q.items = q.items[:lastIndex]

	return item
}
