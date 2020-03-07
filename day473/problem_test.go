package day473

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	g, reversed AdjacencyMatrix
}{
	{
		AdjacencyMatrix{
			'A': {
				'B': struct{}{},
			},
			'B': {
				'C': struct{}{},
			},
		},
		AdjacencyMatrix{
			'C': {
				'B': struct{}{},
			},
			'B': {
				'A': struct{}{},
			},
		},
	},
}

func TestReverseDirectedGraph(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if reversed := ReverseDirectedGraph(tc.g); !reflect.DeepEqual(reversed, tc.reversed) {
			t.Errorf("Expected %v, got %v", tc.reversed, reversed)
		}
	}
}

func BenchmarkReverseDirectedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ReverseDirectedGraph(tc.g)
		}
	}
}
