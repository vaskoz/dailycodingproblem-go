package day372

import "math"

// NumberOfDigits returns the number of digits
// for a natural number input.
// Runs in O(1) time and space and doesn't use loops.
// https://en.wikipedia.org/wiki/Natural_number
func NumberOfDigits(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		n = -n
	}

	return int(math.Log10(float64(n))) + 1
}
