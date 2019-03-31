package day221

import "testing"

var testcases = []struct {
	n, expected int
}{
	{-1, 0},
	{0, 0},
	{1, 1},
	{2, 7},
	{3, 8},
	{4, 49},
	{5, 50},
	{6, 56},
	{7, 57},
	{8, 343},
	{9, 344},
	{10, 350},
}

func TestSevenishNumber(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SevenishNumber(tc.n); result != tc.expected {
			t.Errorf("For n=%d, expected %d, got %d", tc.n, tc.expected, result)
		}
	}
}

func BenchmarkSevenishNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SevenishNumber(tc.n)
		}
	}
}
