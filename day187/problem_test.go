package day187

import "testing"

var testcases = []struct {
	rects    []Rectangle
	expected bool
}{
	{[]Rectangle{{1, 4, 3, 3}, {-1, 3, 2, 1}, {0, 5, 4, 3}}, true},
	{[]Rectangle{{1, 4, 1, 1}, {-1, 3, 1, 1}, {0, 5, 4, 3}}, false},
	{[]Rectangle{{1, 1, 1, 1}, {2, 2, 1, 1}, {3, 3, 1, 1}}, false},
	{[]Rectangle{{1, 1, 2, 1}, {2, 2, 1, 2}, {3, 3, 1, 1}}, true},
}

func TestDoesRectanglePairOverlap(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := DoesRectanglePairOverlap(tc.rects); result != tc.expected {
			t.Errorf("Given %v, expected %v, got %v", tc.rects, tc.expected, result)
		}
	}
}

func BenchmarkDoesRectanglePairOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DoesRectanglePairOverlap(tc.rects)
		}
	}
}
