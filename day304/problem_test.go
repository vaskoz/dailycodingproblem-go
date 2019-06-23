package day304

import "testing"

// nolint
var testcases = []struct {
	row, col, k int
	expected    float64
}{
	{0, 0, 3, 0.125},
	{3, 3, 1, 1.0},
	{3, 3, 2, 0.875},
}

func TestProbKnightOnBoard(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ProbKnightOnBoard(tc.row, tc.col, tc.k); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkProbKnightOnBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ProbKnightOnBoard(tc.row, tc.col, tc.k)
		}
	}
}
