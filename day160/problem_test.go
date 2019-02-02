package day160

import "testing"

var testcases = []struct {
	weights     []WeightedEdges
	longestPath int
}{
	{[]WeightedEdges{
		{"a", "b", 3},
		{"a", "c", 5},
		{"a", "d", 8},
		{"d", "e", 2},
		{"d", "f", 4},
		{"e", "g", 1},
		{"e", "h", 1},
	}, 17},
	{[]WeightedEdges{}, 0},
	{nil, 0},
}

func TestLongestPath(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if path := LongestPath(tc.weights); path != tc.longestPath {
			t.Errorf("Expected %v got %v", tc.longestPath, path)
		}
	}
}

func BenchmarkLongestPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestPath(tc.weights)
		}
	}
}
