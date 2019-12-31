package day407

import (
	"reflect"
	"testing"

	"github.com/algds/kruskals"
)

// nolint
var testcases = []struct {
	pipes    []kruskals.WeightedEdge
	expected []kruskals.WeightedEdge
}{
	{
		[]kruskals.WeightedEdge{
			kruskals.SimpleWeightedEdge{F: 0, T: 1, W: 1},
			kruskals.SimpleWeightedEdge{F: 0, T: 2, W: 5},
			kruskals.SimpleWeightedEdge{F: 0, T: 3, W: 20},
			kruskals.SimpleWeightedEdge{F: 1, T: 3, W: 15},
			kruskals.SimpleWeightedEdge{F: 2, T: 3, W: 10},
		},
		[]kruskals.WeightedEdge{
			kruskals.SimpleWeightedEdge{F: 0, T: 1, W: 1},
			kruskals.SimpleWeightedEdge{F: 0, T: 2, W: 5},
			kruskals.SimpleWeightedEdge{F: 2, T: 3, W: 10},
		},
	},
}

func TestLowestCostPipeConfiguration(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := LowestCostPipeConfiguration(tc.pipes); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkLowestCostPipeConfiguration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LowestCostPipeConfiguration(tc.pipes)
		}
	}
}
