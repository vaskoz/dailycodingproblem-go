package day190

import "testing"

var testcases = []struct {
	nums   []int
	maxSum int
}{
	{[]int{8, -1, 3, 4}, 15},
	{[]int{-4, 5, 1, 0}, 6},
}

func TestMaxWrapSubarraySum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxWrapSubarraySum(tc.nums); result != tc.maxSum {
			t.Errorf("Expected %v, got %v", tc.maxSum, result)
		}
	}
}

func BenchmarkMaxWrapSubarraySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxWrapSubarraySum(tc.nums)
		}
	}
}
