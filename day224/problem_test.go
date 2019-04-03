package day224

import "testing"

var testcases = []struct {
	nums     []int
	smallest int
}{
	{[]int{1, 2, 3, 10}, 7},
	{[]int{1, 3, 4, 5}, 2},
	{[]int{1, 2, 6, 10, 11, 15}, 4},
	{[]int{1, 1, 1, 1}, 5},
	{[]int{1, 1, 3, 4}, 10},
}

func TestSmallestPositiveIntNotInSubsetSum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallestPositiveIntNotInSubsetSum(tc.nums); result != tc.smallest {
			t.Errorf("For %v, expected %v, got %v", tc.nums, tc.smallest, result)
		}
	}
}

func BenchmarkSmallestPositiveIntNotInSubsetSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestPositiveIntNotInSubsetSum(tc.nums)
		}
	}
}
