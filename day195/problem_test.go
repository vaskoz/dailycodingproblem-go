package day195

import "testing"

var testcases = []struct {
	mat            Matrix
	i1, j1, i2, j2 int
	count          int
}{
	{Matrix{
		{1, 3, 7, 10, 15, 20},
		{2, 6, 9, 14, 22, 25},
		{3, 8, 10, 15, 25, 30},
		{10, 11, 12, 23, 30, 35},
		{20, 25, 30, 35, 40, 45},
	},
		1, 1, 3, 3,
		14},
}

func TestCountLargerAndSmallerBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountLargerAndSmallerBrute(tc.mat, tc.i1, tc.j1, tc.i2, tc.j2); result != tc.count {
			t.Errorf("Expected %v, got %v", tc.count, result)
		}
	}
}

func BenchmarkCountLargerAndSmallerBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountLargerAndSmallerBrute(tc.mat, tc.i1, tc.j1, tc.i2, tc.j2)
		}
	}
}

func TestCountLargerAndSmallerEfficient(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountLargerAndSmallerEfficient(tc.mat, tc.i1, tc.j1, tc.i2, tc.j2); result != tc.count {
			t.Errorf("Expected %v, got %v", tc.count, result)
		}
	}
}

func BenchmarkCountLargerAndSmallerEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountLargerAndSmallerEfficient(tc.mat, tc.i1, tc.j1, tc.i2, tc.j2)
		}
	}
}
