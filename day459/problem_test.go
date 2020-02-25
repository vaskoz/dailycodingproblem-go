package day459

import "testing"

// nolint
var testcases = []struct {
	n        int
	expected int
}{
	{4, 1},
	{17, 2},
	{18, 2},
}

func TestSmallestNumberOfPerfectSquaresSum(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := SmallestNumberOfPerfectSquaresSum(tc.n); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkSmallestNumberOfPerfectSquaresSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestNumberOfPerfectSquaresSum(tc.n)
		}
	}
}

func TestSmallestNumberOfPerfectSquaresSumFaster(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := SmallestNumberOfPerfectSquaresSumFaster(tc.n); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkSmallestNumberOfPerfectSquaresSumFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestNumberOfPerfectSquaresSumFaster(tc.n)
		}
	}
}
