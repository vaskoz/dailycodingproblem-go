package day53

import "errors"

var errEmpty = errors.New("empty")

// ErrEmpty is the error returned for an empty data structure.
func ErrEmpty() error {
	return errEmpty
}

// Stack represents a LIFO data structure.
type Stack interface {
	Push(val interface{})
	Pop() (interface{}, error)
}

// Queue represents a FIFO data structure.
type Queue interface {
	Enqueue(item interface{})
	Dequeue() (interface{}, error)
}

type stack struct {
	values []interface{}
}

func (s *stack) Push(val interface{}) {
	s.values = append(s.values, val)
}

func (s *stack) Pop() (interface{}, error) {
	if len(s.values) == 0 {
		return nil, ErrEmpty()
	}
	var x interface{}
	x, s.values = s.values[len(s.values)-1], s.values[:len(s.values)-1]
	return x, nil
}

type twoStacksQueue struct {
	enqueue, dequeue *stack
}

func (tsq twoStacksQueue) Enqueue(item interface{}) {
	tsq.enqueue.Push(item)
}

func (tsq twoStacksQueue) Dequeue() (interface{}, error) {
	if val, err := tsq.dequeue.Pop(); err == nil {
		return val, nil
	}
	for val, err := tsq.enqueue.Pop(); err == nil; val, err = tsq.enqueue.Pop() {
		tsq.dequeue.Push(val)
	}
	if val, err := tsq.dequeue.Pop(); err == nil {
		return val, nil
	}
	return nil, ErrEmpty()
}

// NewQueue returns a queue. FIFO.
func NewQueue() Queue {
	return &twoStacksQueue{&stack{}, &stack{}}
}
