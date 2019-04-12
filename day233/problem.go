package day233

import "errors"

var errOverflow = errors.New("overflow occurred")

// ErrOverflow represents an overflow error.
func ErrOverflow() error {
	return errOverflow
}

// FastFibonnaci runs in O(N) time with O(1) space.
// Returns an error if overflow occurs.
func FastFibonnaci(n uint) (uint, error) {
	if n < 2 {
		return n, nil
	}
	first, second := uint(0), uint(1)
	var err error
	for i := uint(1); i < n; i++ {
		first, second = second, first+second
		if second < first {
			return 0, errOverflow
		}
	}
	return second, err
}
