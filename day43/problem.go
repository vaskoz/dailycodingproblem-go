package day43

import "fmt"

var stackEmptyError = fmt.Errorf("Stack is empty")

// maxIntStack is implemented using O(N) extra memory.
type maxIntStack struct {
	stack []int
}

// MaxIntStack represents that API of a stack of ints that push, pop and find the max value.
type MaxIntStack interface {
	Push(val int)
	Pop() (int, error)
	Max() (int, error)
}

// NewMaxIntStack returns a new instance.
func NewMaxIntStack() MaxIntStack {
	return &maxIntStack{}
}

func (mis *maxIntStack) Push(val int) {
	var max int
	if len(mis.stack) == 0 {
		max = val
	} else {
		max = mis.stack[len(mis.stack)-2]
		if val > max {
			max = val
		}
	}
	mis.stack = append(mis.stack, max, val)
}

func (mis *maxIntStack) Pop() (int, error) {
	if len(mis.stack) == 0 {
		return 0, StackEmptyError()
	}
	var result int
	result, mis.stack = mis.stack[len(mis.stack)-1], mis.stack[:len(mis.stack)-2]
	return result, nil
}

func (mis *maxIntStack) Max() (int, error) {
	if len(mis.stack) == 0 {
		return 0, StackEmptyError()
	}
	return mis.stack[len(mis.stack)-2], nil
}

// StackEmptyError returns the instance of the error representing an empty stack.
func StackEmptyError() error {
	return stackEmptyError
}
