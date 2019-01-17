package day144

import "testing"

var testcases = []struct {
	nums        []int
	targetIndex int
	expected    int
}{
	{[]int{4, 1, 3, 5, 6}, 0, 3},
	{[]int{4, 1, 3, 5, 6}, 4, -1},
	{[]int{4, 1, 3, 5, 6}, 1, 0},
}

func TestClosestLargerNumberIndexBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ClosestLargerNumberIndexBrute(tc.nums, tc.targetIndex); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkClosestLargerNumberIndexBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ClosestLargerNumberIndexBrute(tc.nums, tc.targetIndex)
		}
	}
}
