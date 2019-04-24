package day245

import "testing"

// nolint
var testcases = []struct {
	maxSteps      []int
	expectedJumps int
	impassable    error
}{
	{[]int{6, 2, 4, 0, 5, 1, 1, 4, 2, 9}, 2, nil},
	{[]int{9}, 0, nil},
	{[]int{}, 0, nil},
	{nil, 0, nil},
	{[]int{1, 1, 1, 1, 1, 1, 1}, 6, nil},
	{[]int{1, 1, 0, 1, 1}, 0, ErrImpassable()},
	{[]int{1, 2, 0, 1, 1}, 3, nil},
}

func TestMinimumJumps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if jumps, impassable := MinimumJumps(tc.maxSteps); impassable != tc.impassable || jumps != tc.expectedJumps {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expectedJumps, tc.impassable, jumps, impassable)
		}
	}
}

func BenchmarkMinimumJumps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumJumps(tc.maxSteps) //nolint
		}
	}
}
