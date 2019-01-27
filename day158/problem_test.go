package day158

import "testing"

var testcases = []struct {
	maze     Maze
	expected int
}{
	{[][]int{
		{0, 0, 1},
		{0, 0, 1},
		{1, 0, 0},
	}, 2},
}

func TestNumRoutesThruMaze(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := NumRoutesThruMaze(tc.maze); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkNumRoutesThruMaze(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumRoutesThruMaze(tc.maze)
		}
	}
}
