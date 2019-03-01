package day191

import "testing"

var testcases = []struct {
	intervals []Interval
	removes   int
}{
	{[]Interval{{7, 9}, {2, 4}, {5, 8}}, 1},
}

func TestMinRemoveNoOverlap(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinRemoveNoOverlap(tc.intervals); result != tc.removes {
			t.Errorf("Expected %v, got %v", tc.removes, result)
		}
	}
}

func BenchmarkMinRemoveNoOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinRemoveNoOverlap(tc.intervals)
		}
	}
}
