package day56

import "testing"

var testcases = []struct {
	g        Graph
	k        int
	canColor bool
}{
	{map[int]map[int]struct{}{
		0: map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
		},
		1: map[int]struct{}{
			0: struct{}{},
			2: struct{}{},
		},
		2: map[int]struct{}{
			0: struct{}{},
			1: struct{}{},
			3: struct{}{},
		},
		3: map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
		},
	}, 3, true},
	{map[int]map[int]struct{}{
		0: map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
		},
		1: map[int]struct{}{
			0: struct{}{},
			2: struct{}{},
		},
		2: map[int]struct{}{
			0: struct{}{},
			1: struct{}{},
			3: struct{}{},
		},
		3: map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
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
