package day322

import "testing"

// nolint
var testcases = []struct {
	n, jumps int
}{
	{9, 5},
	{1, 1},
	{0, 0},
	{2, 3},
	{3, 2},
	{-3, 2},
	{-2, 3},
	{-1, 1},
	{-9, 5},
}

func TestNumberOfJumps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if jumps := NumberOfJumps(tc.n); jumps != tc.jumps {
			t.Errorf("Given %v, expected %v, got %v", tc.n, tc.jumps, jumps)
		}
	}
}

func BenchmarkNumberOfJumps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumberOfJumps(tc.n)
		}
	}
}
