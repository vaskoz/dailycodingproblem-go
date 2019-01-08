package day137

import (
	"errors"
)

// BitArray is a bit array.
type BitArray struct {
	size int
	data []uint64
}

var errOutOfBounds = errors.New("Out of bounds of the BitArray")

// ErrOutOfBounds is the error returned if you attempt to access
// an index out of bounds.
func ErrOutOfBounds() error {
	return errOutOfBounds
}

// Init resets the bit array with a new capacity.
// This destroys any previous data.
func (ba *BitArray) Init(size int) {
	ba.size = size
	ba.data = make([]uint64, (size/64)+1)
}

// Set flips the bit at index i to 1-if set=true, and 0-if set=false.
func (ba *BitArray) Set(i int, set bool) error {
	if i < 0 || i >= ba.size {
		return ErrOutOfBounds()
	}
	index := i / 64
	pos := i % 64
	if set {
		mask := uint64(1) << uint(pos)
		ba.data[index] |= mask
	} else {
		mask := uint64(1<<64-1) ^ (uint64(1) << uint(pos))
		ba.data[index] &= mask
	}
	return nil
}

// Get returns the value of the bit.
// TRUE=1, FALSE, 0
func (ba *BitArray) Get(i int) (bool, error) {
	if i < 0 || i >= ba.size {
		return false, ErrOutOfBounds()
	}
	index := i / 64
	pos := i % 64
	mask := uint64(1) << uint(pos)
	return (ba.data[index] & mask) != 0, nil
}
