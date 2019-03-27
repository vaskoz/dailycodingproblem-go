package day72

import "testing"

var testcases = []struct {
	nodes       []Node
	edges       []Edge
	expected    int
	expectedErr error
}{
	{[]Node("ABACA"), []Edge{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3, nil},
	{[]Node("ABAACB"), []Edge{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {4, 5}}, 2, nil},
	{[]Node("A"), []Edge{{0, 0}}, 0, ErrInfiniteLoop()},
	{[]Node("AAAA"), []Edge{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, 0, ErrInfiniteLoop()},
}

func TestLargestPathValue(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := LargestPathValue(tc.nodes, tc.edges); result != tc.expected || err != tc.expectedErr {
			t.Errorf("Expected (%v,%v) but got (%v,%v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func BenchmarkLargestPathValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestPathValue(tc.nodes, tc.edges)
		}
	}
}
