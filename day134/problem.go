package day134

import "fmt"

// SparseArray is a space-efficient array for one containing
// mostly zeros.
type SparseArray struct {
	data map[int]int
	size int
}

var errIndexOutOfRange = fmt.Errorf("index is out of range")

// ErrIndexOutOfRange returns the error that's associated with access beyond
// the limits of the SparseArray.
func ErrIndexOutOfRange() error {
	return errIndexOutOfRange
}

// NewSparseArray initializes a SparseArray with an original array.
func NewSparseArray(arr []int) *SparseArray {
	sa := &SparseArray{make(map[int]int), len(arr)}
	for i := range arr {
		if arr[i] != 0 {
			sa.data[i] = arr[i]
		}
	}
	return sa
}

// Set mutates the value at position 'i' with value 'val'.
func (sa *SparseArray) Set(i, val int) error {
	if sa == nil {
		return fmt.Errorf("create with NewSparseArray")
	}
	if i >= sa.size || i < 0 {
		return errIndexOutOfRange
	}
	if val == 0 {
		delete(sa.data, i)
	} else {
		sa.data[i] = val
	}
	return nil
}

// Get returns the value at position 'i'.
func (sa *SparseArray) Get(i int) (int, error) {
	if sa == nil {
		return 0, fmt.Errorf("create with NewSparseArray")
	}
	if i >= sa.size || i < 0 {
		return 0, errIndexOutOfRange
	}
	return sa.data[i], nil
}
