package day428

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	stones, pyramid []int
	cost            int
}{
	{
		[]int{1, 1, 3, 3, 2, 1},
		[]int{0, 1, 2, 3, 2, 1},
		2,
	},
	{
		[]int{0, 0, 0, 9, 9, 0},
		[]int{0, 0, 0, 1, 0, 0},
		17,
	},
	{
		[]int{2, 9, 7, 1},
		[]int{1, 2, 1, 0},
		15,
	},
	{
		[]int{2, 9, 7, 9, 9},
		[]int{1, 2, 3, 2, 1},
		27,
	},
	{
		[]int{9, 9, 9},
		[]int{1, 2, 1},
		23,
	},
}

func TestLowestCostPyramid(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if pyramid, cost := LowestCostPyramid(tc.stones); cost != tc.cost ||
			!reflect.DeepEqual(tc.pyramid, pyramid) {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.pyramid, tc.cost, pyramid, cost)
		}
	}
}

func BenchmarkLowestCostPyramid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LowestCostPyramid(tc.stones)
		}
	}
}
