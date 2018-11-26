package day95

import "sort"

// NextPermutation reorders the digits for the next
// greater permutation.
func NextPermutation(digits []int) {
	var i int
	for i = len(digits) - 1; i > 0; i-- {
		if digits[i] > digits[i-1] {
			break
		}
	}
	if i == 0 {
		sort.Ints(digits)
		return
	}
	x := digits[i-1]
	min := i
	for j := i + 1; j < len(digits); j++ {
		if digits[j] > x && digits[j] < digits[min] {
			min = j
		}
	}
	digits[i-1], digits[min] = digits[min], digits[i-1]
	sort.Ints(digits[i:])
}
