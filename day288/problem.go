package day288

import (
	"errors"
	"reflect"
	"sort"
)

var errInvalidInput = errors.New("only 4 digit positive numbers allowed")

// ErrInvalidInput returns the error associated with invalid input.
func ErrInvalidInput() error {
	return errInvalidInput
}

// KaprekarConstantIterations returns the number of iterations required to
// reach Kaprekar's constant using Kaprekar's routine.
func KaprekarConstantIterations(num int) (int, error) {
	if num < 1 || num > 9999 {
		return 0, ErrInvalidInput()
	}
	digits := convertIntToSlice(num)
	iterations := 0
	kaprekarsConstant := []int{6, 1, 7, 4}
	last := make([]int, 4)
	for !reflect.DeepEqual(digits, kaprekarsConstant) {
		if reflect.DeepEqual(last, digits) {
			return 0, ErrInvalidInput()
		}
		desc := make([]int, 4)
		asc := make([]int, 4)
		copy(desc, digits)
		copy(asc, digits)
		sort.Sort(sort.Reverse(sort.IntSlice(desc)))
		sort.Ints(asc)
		result := convertSliceToInt(desc) - convertSliceToInt(asc)
		digits = convertIntToSlice(result)
		iterations++
	}
	return iterations, nil
}

func convertSliceToInt(digits []int) int {
	pow := 1000
	var result int
	for _, digit := range digits {
		result += digit * pow
		pow /= 10
	}
	return result
}

func convertIntToSlice(num int) []int {
	digits := make([]int, 4)
	pos := 3
	for num != 0 {
		digits[pos] = num % 10
		num /= 10
		pos--
	}
	return digits
}
