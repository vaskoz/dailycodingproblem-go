package day136

import "testing"

var testcases = []struct {
	matrix  [][]int
	maxArea int
}{
	{[][]int{
		{1, 0, 0, 0},
		{1, 0, 1, 1},
		{1, 0, 1, 1},
		{0, 1, 0, 0},
	}, 4},
	{[][]int{
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}, 4},
	{[][]int{
		{1, 1, 1, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
	}, 3},
	{[][]int{
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}, 8},
}

func TestMaxRectangleOfOnesAreaBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxRectangleOfOnesAreaBrute(tc.matrix); result != tc.maxArea {
			t.Errorf("Expected %v got %v", tc.maxArea, result)
		}
	}
}

func BenchmarkMaxRectangleOfOnesAreaBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxRectangleOfOnesAreaBrute(tc.matrix)
		}
	}
}
