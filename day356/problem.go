package day356

import "errors"

const fixedArraySize = 10

var errEmpty = errors.New("queue is empty")
var errFull = errors.New("queue is full")

// ErrEmpty is the error returned if queue is empty.
func ErrEmpty() error {
	return errEmpty
}

// ErrFull is the error returned if queue is full.
func ErrFull() error {
	return errFull
}

// NewQueueFLA returns a new instance of QueueFLA
// with a given capacity.
func NewQueueFLA(capacity int) QueueFLA {
	slots := capacity / fixedArraySize
	if capacity%fixedArraySize != 0 {
		slots++
	}
	return &queueFLA{
		set:       make([][fixedArraySize]interface{}, slots),
		size:      0,
		capacity:  capacity,
		frontSlot: -1,
		frontPos:  -1,
		backSlot:  -1,
		backPos:   -1,
	}
}

// QueueFLA is a queue implemented using a set
// of fixed-length arrays.
// In this implementation, the fixed array size is 10.
type QueueFLA interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	Size() int
}

type queueFLA struct {
	set                 [][fixedArraySize]interface{}
	size, capacity      int
	frontSlot, frontPos int
	backSlot, backPos   int
}

// Enqueue attempts to add an item to the queue if there is room.
func (q *queueFLA) Enqueue(item interface{}) error {
	if q.size == q.capacity {
		return errFull
	}
	q.size++
	if q.frontSlot == -1 && q.frontPos == -1 {
		q.frontSlot++
		q.frontPos++
		q.set[q.frontSlot][q.frontPos] = item
		q.backSlot++
		q.backPos++
	} else {
		q.backPos++
		if q.backPos == fixedArraySize {
			q.backPos = 0
			q.backSlot++
			if q.backSlot == len(q.set) {
				q.backSlot = 0
			}
		}
		q.set[q.backSlot][q.backPos] = item
	}
	return nil
}

// Dequeue attempts to remove an item from the queue if
// the queue is not empty.
func (q *queueFLA) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, errEmpty
	}
	q.size--
	res := q.set[q.frontSlot][q.frontPos]
	q.frontPos++
	if q.frontPos == fixedArraySize {
		q.frontPos = 0
		q.frontSlot++
		if q.frontSlot == len(q.set) {
			q.frontSlot = 0
		}
	}
	return res, nil
}

// Size returns the number of elements currently in the queue.
func (q *queueFLA) Size() int {
	return q.size
}
