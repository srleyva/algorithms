package algorithms

// Queue is a Queue implemtation in go FIFO
type Queue struct {
	queue []int
}

// NewQueue returns a new Queue
func NewQueue() *Queue {
	queue := make([]int, 0)
	return &Queue{queue}
}

// Top will show the value of the top element without removing it
func (q *Queue) Top() int {
	return q.queue[0]
}

// Pop takes the first element off the queue and removes
func (q *Queue) Pop() int {
	rval := q.Top()
	q.queue = q.queue[1:]
	return rval
}

// Push adds an item to the queue
func (q *Queue) Push(t int) {
	q.queue = append(q.queue, t)
}

// IsEmpty will check if the queue is empty and return true if so
func (q *Queue) IsEmpty() bool {
	if len(q.queue) == 0 {
		return true
	}
	return false
}
