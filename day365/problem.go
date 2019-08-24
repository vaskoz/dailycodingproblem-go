package day365

import "errors"

var errEmptyStack = errors.New("stack is empty")

// ErrEmptyStack is the returned error for an empty stack.
func ErrEmptyStack() error {
	return errEmptyStack
}

// Quack is a data structure combining properties of both
// stacks and queues.
type Quack struct {
	left, right, buffer []interface{}
}

// Push adds a new item x to the left end of the list.
func (q *Quack) Push(x interface{}) {
	q.left = append(q.left, x)
}

func (q *Quack) common(a, b *[]interface{}) (interface{}, error) {
	if len(*a) == 0 && len(*b) == 0 {
		return nil, errEmptyStack
	}
	if len(*a) == 0 {
		for i := 0; i < len(*b)/2; i++ {
			last := len(*b) - 1
			x := (*b)[last]
			*b = (*b)[:last]
			q.buffer = append(q.buffer, x)
		}
		for len(*b) != 0 {
			last := len(*b) - 1
			x := (*b)[last]
			*b = (*b)[:last]
			*a = append(*a, x)
		}
		for len(q.buffer) != 0 {
			last := len(q.buffer) - 1
			x := q.buffer[last]
			q.buffer = q.buffer[:last]
			*b = append(*b, x)
		}
	}
	last := len(*a) - 1
	res := (*a)[last]
	*a = (*a)[:last]
	return res, nil
}

// Pop removes and return the item on the left end of the list.
func (q *Quack) Pop() (interface{}, error) {
	return q.common(&q.left, &q.right)
}

// Pull removes the item on the right end of the list.
func (q *Quack) Pull() (interface{}, error) {
	return q.common(&q.right, &q.left)
}
