package day217

import "testing"

var testcases = []struct {
	n, expected uint
}{
	{21, 21},
	{22, 32},
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 4},
}

func TestSmallestSparseNumberGreaterThan(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallestSparseNumberGreaterThan(tc.n); result != tc.expected {
			t.Errorf("For N=%v, expected %v, got %v", tc.n, tc.expected, result)
		}
	}
}

func BenchmarkSmallestSparseNumberGreaterThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestSparseNumberGreaterThan(tc.n)
		}
	}
}
