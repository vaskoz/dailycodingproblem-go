package day245

import "testing"

var testcases = []struct {
	maxSteps      []int
	expectedJumps int
}{
	{[]int{6, 2, 4, 0, 5, 1, 1, 4, 2, 9}, 2},
	{[]int{9}, 0},
	{[]int{}, 0},
	{nil, 0},
	{[]int{1, 1, 1, 1, 1, 1, 1}, 6},
}

func TestMinimumJumps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if jumps := MinimumJumps(tc.maxSteps); jumps != tc.expectedJumps {
			t.Errorf("Expected %v, got %v", tc.expectedJumps, jumps)
		}
	}
}

func BenchmarkMinimumJumps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumJumps(tc.maxSteps)
		}
	}
}
