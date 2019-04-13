package day234

import (
	"reflect"
	"testing"

	"github.com/algds/kruskals"
)

var testcases = []struct {
	graph, maximumMST []kruskals.SimpleWeightedEdge
}{
	{
		[]kruskals.SimpleWeightedEdge{{0, 3, 3}, {3, 1, 30}, {0, 1, 20}, {0, 4, 10}, {1, 4, 5}, {4, 2, 20}, {1, 2, 50}},
		[]kruskals.SimpleWeightedEdge{{1, 2, 50}, {3, 1, 30}, {0, 1, 20}, {4, 2, 20}},
	},
}

func TestMaximumSpanningTree(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaximumSpanningTree(tc.graph); !reflect.DeepEqual(result, tc.maximumMST) {
			t.Errorf("Expected %v, got %v", tc.maximumMST, result)
		}
	}
}

func BenchmarkMaximumSpanningTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaximumSpanningTree(tc.graph)
		}
	}
}
