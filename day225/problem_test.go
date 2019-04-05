package day225

import "testing"

var testcases = []struct {
	n, k, survivor int
}{
	{5, 2, 3},
	{41, 2, 19},
	{5, 3, 4},
	{1, 100, 1},
	{-20, 100, 0},
	{0, 100, 0},
}

func TestJosephusGeneral(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if survivor := JosephusGeneral(tc.n, tc.k); survivor != tc.survivor {
			t.Errorf("Given N=%d, k=%d, expected %d, got %d", tc.n, tc.k, tc.survivor, survivor)
		}
	}
}

func BenchmarkJosephusGeneral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			JosephusGeneral(tc.n, tc.k)
		}
	}
}
