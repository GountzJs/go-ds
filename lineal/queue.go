package lineal

type queue struct {
	items    []int
	capacity uint
	head     int
	tail     int
}

func NewQueue(capacity uint) *queue {
	return &queue{
		capacity: capacity,
		items:    make([]int, capacity),
		head:     -1,
		tail:     -1,
	}
}

func (q queue) IsEmpty() bool {
	return q.head == -1 && q.tail == -1
}

func (q queue) IsFull() bool {
	return q.tail == int(q.capacity-1)
}

func (q *queue) Enqueue(value int) {
	if q.IsFull() {
		return
	}

	if q.head == -1 {
		var initialIndex = 0
		q.head = initialIndex
		q.tail = initialIndex
		q.items[initialIndex] = value
		return
	}

	q.tail = q.tail + 1
	q.items[q.tail] = value
}

func (q *queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	if q.head == q.tail {
		var resetIndex = -1
		var valueOfIndex = q.items[q.head]
		q.head = resetIndex
		q.tail = resetIndex
		return valueOfIndex
	}

	var valueOfIndex = q.items[q.head]
	q.head = q.head + 1

	return valueOfIndex
}

func (q queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.items[q.head]
}
