package day257

import "testing"

// nolint
var testcases = []struct {
	nums       []int
	start, end int
}{
	{[]int{3, 7, 5, 6, 9}, 1, 3},
}

func TestSmallestWindowThatMustBeSorted(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if start, end := SmallestWindowThatMustBeSorted(tc.nums); start != tc.start || end != tc.end {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.start, tc.end, start, end)
		}
	}
}

func BenchmarkSmallestWindowThatMustBeSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestWindowThatMustBeSorted(tc.nums)
		}
	}
}
