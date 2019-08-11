package day353

import "testing"

// nolint
var testcases = []struct {
	hist []int
	area int
}{
	{[]int{1, 3, 2, 5}, 6},
	{[]int{6, 2, 5, 4, 5, 1, 6}, 12},
}

func TestLargestRectangleHistogram(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if area := LargestRectangleHistogram(tc.hist); area != tc.area {
			t.Errorf("Expected %v, got %v", tc.area, area)
		}
	}
}

func BenchmarkLargestRectangleHistogram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestRectangleHistogram(tc.hist)
		}
	}
}
