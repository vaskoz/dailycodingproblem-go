package day166

import (
	"errors"
)

var errNoMoreElements = errors.New("no more elements to iterate")

// ErrNoMoreElements is returned when there are no more elements to iterate.
func ErrNoMoreElements() error {
	return errNoMoreElements
}

// TwoDimensionalIterator can iterate over a slice of slices of anything.
type TwoDimensionalIterator interface {
	Next() (interface{}, error)
	HasNext() bool
}

type twoDimensionalIterator struct {
	sliceIndex, positionInSlice int
	finished                    bool
	data                        [][]interface{}
}

// NewTwoDimensionalIterator returns a new instance of TwoDimensionalIterator.
func NewTwoDimensionalIterator(data [][]interface{}) TwoDimensionalIterator {
	finished := true
	var sliceIndex, positionInSlice int
	for i := range data {
		for j := range data[i] {
			finished = false
			sliceIndex = i
			positionInSlice = j
			break
		}
		if !finished {
			break
		}
	}
	return &twoDimensionalIterator{
		sliceIndex: sliceIndex, finished: finished, positionInSlice: positionInSlice, data: data,
	}
}

func (tdi *twoDimensionalIterator) Next() (interface{}, error) {
	if tdi.finished {
		return nil, ErrNoMoreElements()
	}
	sliceIndex := -1
	positionIndex := 0
	for i := tdi.sliceIndex; i < len(tdi.data); i++ {
		positionIndex = 0
		if i == tdi.sliceIndex {
			positionIndex = tdi.positionInSlice + 1
		}
		for j := positionIndex; j < len(tdi.data[i]); j++ {
			sliceIndex = i
			positionIndex = j
			break
		}
		if sliceIndex != -1 {
			break
		}
	}
	result := tdi.data[tdi.sliceIndex][tdi.positionInSlice]
	if sliceIndex == -1 {
		tdi.finished = true
	} else {
		tdi.sliceIndex = sliceIndex
		tdi.positionInSlice = positionIndex
	}
	return result, nil
}

func (tdi *twoDimensionalIterator) HasNext() bool {
	return !tdi.finished
}
