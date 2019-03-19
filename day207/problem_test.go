package day207

import "testing"

var testcases = []struct {
	g         UndirectedGraph
	isPartite bool
}{
	{UndirectedGraph{
		"A": {
			"B": {},
			"D": {},
		},
		"B": {
			"A": {},
			"C": {},
		},
		"C": {
			"B": {},
			"D": {},
		},
		"D": {
			"A": {},
			"C": {},
		},
	}, true},
	{UndirectedGraph{
		"A": {
			"B": {},
			"C": {},
		},
		"B": {
			"A": {},
			"C": {},
		},
		"C": {
			"A": {},
			"B": {},
		},
	}, false},
}

func TestIsBipartiteGraph(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if isPartite := IsBipartiteGraph(tc.g); isPartite != tc.isPartite {
			t.Errorf("Expected %v, got %v", tc.isPartite, isPartite)
		}
	}
}

func BenchmarkIsBipartiteGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsBipartiteGraph(tc.g)
		}
	}
}
