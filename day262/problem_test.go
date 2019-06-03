package day262

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	graph   []Edge
	bridges []Edge
}{
	{
		[]Edge{{1, 0}, {0, 2}, {2, 1}, {0, 3}, {3, 4}},
		[]Edge{{3, 4}, {0, 3}},
	},
	{
		[]Edge{{0, 1}, {1, 2}, {2, 3}},
		[]Edge{{2, 3}, {1, 2}, {0, 1}},
	},
	{
		[]Edge{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {1, 4}, {1, 6}, {3, 5}, {4, 5}},
		[]Edge{{1, 6}},
	},
}

func TestFindAllBridges(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if bridges := FindAllBridges(tc.graph); !reflect.DeepEqual(bridges, tc.bridges) {
			t.Errorf("Expected %v, got %v", tc.bridges, bridges)
		}
	}
}

func BenchmarkFindAllBridges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindAllBridges(tc.graph)
		}
	}
}
