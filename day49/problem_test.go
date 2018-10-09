package day49

import "testing"

var testcases = []struct {
	arr    []int
	maxSum int
}{
	{[]int{34, -50, 42, 14, -5, 86}, 137},
	{[]int{-5, -1, -8, -9}, 0},
}

func TestMaxSubArray(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxSubArray(tc.arr); result != tc.maxSum {
			t.Errorf("Expected %d got %d", tc.maxSum, result)
		}
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxSubArray(tc.arr)
		}
	}
}
