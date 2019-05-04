package day255

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	g        Graph
	expected [][]int
}{
	{
		Graph{
			0: {0, 1, 3},
			1: {1, 2},
			2: {2},
			3: {3},
		},
		[][]int{
			{1, 1, 1, 1},
			{0, 1, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	},
	{
		Graph{
			0: {1, 3},
			1: {2},
			2: {},
			3: {},
		},
		[][]int{
			{0, 1, 1, 1},
			{0, 0, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	},
	{
		Graph{
			0: {1, 3},
			1: {2},
			2: {3},
			3: {2},
		},
		[][]int{
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
			{0, 0, 1, 0},
		},
	},
}

func TestTransitiveClosure(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := TransitiveClosure(tc.g); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkTransitiveClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TransitiveClosure(tc.g)
		}
	}
}
