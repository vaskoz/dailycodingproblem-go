package day281

import "testing"

// nolint
var testcases = []struct {
	bricks [][]int
	fewest int
}{
	{
		[][]int{
			{3, 5, 1, 1},
			{2, 3, 3, 2},
			{5, 5},
			{4, 4, 2},
			{1, 3, 3, 3},
			{1, 1, 6, 1, 1},
		},
		2,
	},
}

func TestFewestBricksCut(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := FewestBricksCut(tc.bricks); result != tc.fewest {
			t.Errorf("Expected %v, got %v", tc.fewest, result)
		}
	}
}

func BenchmarkFewestBricksCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FewestBricksCut(tc.bricks)
		}
	}
}
