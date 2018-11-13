package day84

import "testing"

var testcases = []struct {
	grid     [][]int
	expected int
}{
	{[][]int{
		{1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
	}, 4},
}

func TestCountIslands(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountIslands(tc.grid); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkCountIslands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountIslands(tc.grid)
		}
	}
}
