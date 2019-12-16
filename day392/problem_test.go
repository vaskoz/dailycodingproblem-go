package day392

import "testing"

// nolint
var testcases = []struct {
	grid      [][]int
	perimeter int
}{
	{
		[][]int{
			{0, 1, 1, 0},
			{1, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 1, 0},
		},
		14,
	},
	{
		[][]int{
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
		},
		12,
	},
	{
		[][]int{
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
			{1, 1, 1, 1},
		},
		16,
	},
	{
		[][]int{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
		14,
	},
}

func TestIslandPerimeter(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := IslandPerimeter(tc.grid); res != tc.perimeter {
			t.Errorf("Expected %v, got %v", tc.perimeter, res)
		}
	}
}

func BenchmarkIslandPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IslandPerimeter(tc.grid)
		}
	}
}
