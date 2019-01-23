package day151

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	grid     PixelGrid
	coord    Coordinate
	color    Color
	expected PixelGrid
}{
	{
		PixelGrid{
			[]Color("BBW"),
			[]Color("WWW"),
			[]Color("WWW"),
			[]Color("BBB"),
		},
		Coordinate{2, 2},
		Color('G'),
		PixelGrid{
			[]Color("BBG"),
			[]Color("GGG"),
			[]Color("GGG"),
			[]Color("BBB"),
		},
	},
}

func TestReplaceAdjacentColorPixel(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		testGrid := copyGrid(tc.grid)
		ReplaceAdjacentColorPixel(testGrid, tc.coord, tc.color)
		if !reflect.DeepEqual(tc.expected, testGrid) {
			t.Errorf("Expected %v got %v", tc.expected, testGrid)
		}
	}
}

func BenchmarkReplaceAdjacentColorPixel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			testGrid := copyGrid(tc.grid)
			ReplaceAdjacentColorPixel(testGrid, tc.coord, tc.color)
		}
	}
}

func copyGrid(grid PixelGrid) PixelGrid {
	var newGrid PixelGrid
	for i := range grid {
		newGrid = append(newGrid, append([]Color{}, grid[i]...))
	}
	return newGrid
}
