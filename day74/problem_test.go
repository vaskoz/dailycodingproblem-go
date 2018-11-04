package day74

import "testing"

var testcases = []struct {
	N, X     int
	expected int
}{
	{6, 12, 4},
	{6, 36, 1},
	{6, 37, 0},
	{6, 29, 0},
	{1000, 29, 2},
}

func TestMultiplicationTableBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MultiplicationTableBrute(tc.N, tc.X); result != tc.expected {
			t.Errorf("Given N=%d X=%d got %d expected %d", tc.N, tc.X, result, tc.expected)
		}
	}
}

func BenchmarkMultiplicationTableBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MultiplicationTableBrute(tc.N, tc.X)
		}
	}
}

func TestMultiplicationTableLinear(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MultiplicationTableLinear(tc.N, tc.X); result != tc.expected {
			t.Errorf("Given N=%d X=%d got %d expected %d", tc.N, tc.X, result, tc.expected)
		}
	}
}

func BenchmarkMultiplicationTableLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MultiplicationTableLinear(tc.N, tc.X)
		}
	}
}
