package day436

import (
	"errors"
)

type node struct {
	value interface{}
	prev  *node
}

// ThreeStacks represents 3-stacks
// using only a single underlying list to store values.
type ThreeStacks struct {
	data         []*node
	heads        [3]*node
	pushes, pops int
}

// minDefrag represents the minimum number of Pops before
// a defrag will happen.
const minDefrag = 20

var errInvalidStackID = errors.New("invalid stack id")
var errEmptyStack = errors.New("stack is empty")

// ErrInvalidStackID is returned if the stack id is not either 0, 1, 2.
func ErrInvalidStackID() error {
	return errInvalidStackID
}

// ErrEmptyStack is returned if the stack is empty.
func ErrEmptyStack() error {
	return errEmptyStack
}

// Pop removes and returns the most recent addition to the given
// stack.
func (ts *ThreeStacks) Pop(stackID int) (interface{}, error) {
	if stackID < 0 || stackID >= 3 {
		return nil, errInvalidStackID
	}

	if ts.heads[stackID] == nil {
		return nil, errEmptyStack
	}
	ts.pops++
	top := ts.heads[stackID]
	ts.heads[stackID] = top.prev

	if ts.pops > minDefrag {
		defrag := make([]*node, 0, ts.pushes-ts.pops)
		pushes := 0

		for i := 0; i < 3; i++ {
			ptr := ts.heads[i]
			for ptr != nil {
				defrag = append(defrag, ptr)
				ptr = ptr.prev
				pushes++
			}
		}

		ts.data = defrag
		ts.pops = 0
		ts.pushes = pushes
	}

	return top.value, nil
}

// Push adds the Value to the top of the given stack.
func (ts *ThreeStacks) Push(val interface{}, stackID int) error {
	if stackID < 0 || stackID > 2 {
		return errInvalidStackID
	}
	ts.pushes++

	v := &node{value: val, prev: ts.heads[stackID]}
	ts.data = append(ts.data, v)
	ts.heads[stackID] = v

	return nil
}
