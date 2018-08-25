package day4

import (
	"testing"
)

var testcases = []struct {
	input                   []int
	smallestMissingPositive int
}{
	{[]int{-10, -20, -40, 1, 2, 3, 4, 5}, 6},
	{[]int{3, 4, -1, 1}, 2},
	{[]int{1, 2, 0}, 3},
	{[]int{1, -1, 2, -2, 0, 100, -100}, 3},
	{[]int{-1, -1, -2, -2, 0, -100, -100}, 1},
	{[]int{1, 1, 2, 2, 0, 3, 4, 5, 6, 7, 8, 9, 9, 8, -100}, 10},
	{[]int{100, 99, 100, -20, -30, -20}, 1},
}

func TestFindSmallestMissingPositive(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := FindSmallestMissingPositive(tc.input); result != tc.smallestMissingPositive {
			t.Errorf("Input was %v. I expected %d, but received %d", tc.input, tc.smallestMissingPositive, result)
		}
	}
}

func BenchmarkFindSmallestMissingPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindSmallestMissingPositive(tc.input)
		}
	}
}
