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
		[]kruskals.SimpleWeightedEdge{
			{F: 0, T: 3, W: 3}, {F: 3, T: 1, W: 30}, {F: 0, T: 1, W: 20},
			{F: 0, T: 4, W: 10}, {F: 1, T: 4, W: 5}, {F: 4, T: 2, W: 20},
			{F: 1, T: 2, W: 50}},
		[]kruskals.SimpleWeightedEdge{
			{F: 1, T: 2, W: 50}, {F: 3, T: 1, W: 30},
			{F: 0, T: 1, W: 20}, {F: 4, T: 2, W: 20},
		},
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
