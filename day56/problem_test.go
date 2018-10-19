package day56

import "testing"

var testcases = []struct {
	g        Graph
	k        int
	canColor bool
}{
	{map[int]map[int]struct{}{
		0: {
			1: {},
			2: {},
			3: {},
		},
		1: {
			0: {},
			2: {},
		},
		2: {
			0: {},
			1: {},
			3: {},
		},
		3: {
			1: {},
			2: {},
		},
	}, 3, true},
	{map[int]map[int]struct{}{
		0: {
			1: {},
			2: {},
			3: {},
		},
		1: {
			0: {},
			2: {},
		},
		2: {
			0: {},
			1: {},
			3: {},
		},
		3: {
			1: {},
			2: {},
		},
	}, 2, false},
}

func TestCanColor(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if canColor := CanColor(0, tc.g, make([]int, len(tc.g)), tc.k); canColor != tc.canColor {
			t.Errorf("Expected %v got %v", tc.canColor, canColor)
		}
	}
}

func BenchmarkCanColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CanColor(0, tc.g, make([]int, len(tc.g)), tc.k)
		}
	}
}
