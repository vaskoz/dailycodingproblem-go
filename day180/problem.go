package day180

// Stack is a stack with only push & pop operations.
type Stack struct {
	data []interface{}
}

// Push adds an element to the stack.
func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

// Pop removes and element from the stack.
func (s *Stack) Pop() interface{} {
	var v interface{}
	v, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return v
}

// Len returns the number of items in the stack.
func (s *Stack) Len() int {
	return len(s.data)
}

// Queue is a simple FIFO queue.
type Queue struct {
	data []interface{}
}

// Enqueue adds to the queue.
func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

// Dequeue removes an element from the queue.
func (q *Queue) Dequeue() interface{} {
	var v interface{}
	v, q.data = q.data[0], q.data[1:]
	return v
}

// Len returns the number of items in the queue.
func (q *Queue) Len() int {
	return len(q.data)
}

// InterleaveStack interleaves the first half of the stack
// with the reverse of the second half.
// This modifies the stack in place.
// Uses only a queue for additional help.
func InterleaveStack(s Stack) {
	var q Queue
	for cycles := s.Len() - 1; cycles > 0; cycles-- {
		for i := 0; i < cycles; i++ {
			q.Enqueue(s.Pop())
		}
		for q.Len() != 0 {
			s.Push(q.Dequeue())
		}
	}
}
