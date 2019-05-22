package day273

import "errors"

var errNoFixedPoint = errors.New("no fixed point found")

// ErrNoFixedPoint returns the error associated with no fixed point in the input.
func ErrNoFixedPoint() error {
	return errNoFixedPoint
}

// FindFixedPoint returns the fixed point number/index if it exists.
// Otherwise, it returns an error if it doesn't exist.
// Runs in O(log N) time and O(1) space.
func FindFixedPoint(sorted []int) (int, error) {
	left, right := 0, len(sorted)
	for left < right {
		switch mid := (left + right) / 2; {
		case mid == sorted[mid]:
			return mid, nil
		case mid < sorted[mid]:
			right = mid
		default:
			left = mid
		}
	}
	return 0, errNoFixedPoint
}
