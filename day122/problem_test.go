package day122

import "testing"

var testcases = []struct {
	grid     [][]int
	maxCoins int
}{
	{[][]int{
		{0, 3, 1, 1},
		{2, 0, 0, 4},
		{1, 5, 3, 1}},
		12},
}

func TestMaxCoinsPath(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if coins := MaxCoinsPath(tc.grid); coins != tc.maxCoins {
			t.Errorf("Expected %v got %v", tc.maxCoins, coins)
		}
	}
}

func BenchmarkMaxCoinsPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxCoinsPath(tc.grid)
		}
	}
}
