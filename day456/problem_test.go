package day456

import "testing"

// nolint
var testcases = []struct {
	moves      []Movement
	start, end uint64
}{
	{[]Movement{{0, 3, Enter}, {1, 3, Exit}}, 0, 1},
	{[]Movement{}, 0, 0},
	{nil, 0, 0},
	{[]Movement{{30, 20, Exit}, {10, 10, Enter}, {20, 10, Enter}}, 20, 30},
}

func TestBusiestBuildingTimes(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if start, end := BusiestBuildingTimes(tc.moves); start != tc.start || end != tc.end {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.start, tc.end, start, end)
		}
	}
}

func BenchmarkBusiestBuildingTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BusiestBuildingTimes(tc.moves)
		}
	}
}
