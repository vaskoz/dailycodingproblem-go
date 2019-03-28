package day218

import (
	"reflect"
	"testing"
)

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
