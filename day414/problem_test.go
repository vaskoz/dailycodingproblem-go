package day414

import "testing"

// nolint
var testcases = []struct {
	n, expected int
}{
	{0, 1},
	{1, 1},
	{2, 0},
	{3, 0},
	{4, 2},
	{5, 10},
	{6, 4},
	{7, 40},
	{8, 92},
	{9, 352},
}

func TestNQueens(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := NQueens(tc.n, []int{}); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkNQueens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NQueens(tc.n, []int{})
		}
	}
}
