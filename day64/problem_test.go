package day64

import "testing"

var testcases = []struct {
	N, Tours int
}{
	{1, 1},
	{2, 0},
	{3, 0},
	{4, 0},
	// {5, 1728}, // uncommenting takes 94.08s user 1.04s system 104% cpu 1:31.42 total
}

func TestKnightsTourCountBacktracking(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := KnightsTourCountBacktracking(tc.N); result != tc.Tours {
			t.Errorf("For N=%d expected %d got %d", tc.N, tc.Tours, result)
		}
	}
}

func BenchmarkKnightsTourCountBacktracking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KnightsTourCountBacktracking(3)
	}
}
