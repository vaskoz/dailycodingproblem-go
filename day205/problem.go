package day205

import (
	"errors"
	"sort"
)

var errNoGreaterPermutation = errors.New("greater permutation impossible")

// ErrNoGreaterPermutation is the error when the greater
// permutation doesn't exist.
func ErrNoGreaterPermutation() error {
	return errNoGreaterPermutation
}

// NextPermutationUint find the next greater permutation
// of it's digits in absolute order.
// Return an error if a greater permutation isn't possible.
func NextPermutationUint(num uint) (uint, error) {
	digits := separateDigits(num)
	var i int
	for i = len(digits) - 1; i > 0; i-- {
		if digits[i] > digits[i-1] {
			break
		}
	}
	if i == 0 {
		sort.Ints(digits)
		return 0, ErrNoGreaterPermutation()
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
	return convertToUint(digits), nil
}

func convertToUint(digits []int) uint {
	var result uint
	multiplier := 1
	for i := range digits {
		result += uint(digits[len(digits)-1-i] * multiplier)
		multiplier *= 10
	}
	return result
}

func separateDigits(num uint) []int {
	var result []int
	for num != 0 {
		result = append(result, int(num)%10)
		num /= 10
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
