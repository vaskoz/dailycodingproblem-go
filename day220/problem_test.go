package day220

import "testing"

var testcases = []struct {
	coins     []int
	maxProfit int
}{
	{[]int{6, 9, 1, 2, 16, 8}, 23},
	{[]int{8, 15, 3, 7}, 22},
	{[]int{2, 2, 2, 2}, 4},
	{[]int{20, 30, 2, 2, 2, 10}, 42},
}

func TestAlternatingCoinsFirst(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := AlternatingCoinsFirst(tc.coins); result != tc.maxProfit {
			t.Errorf("Expected %v, got %v", tc.maxProfit, result)
		}
	}
}

func BenchmarkAlternatingCoinsFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AlternatingCoinsFirst(tc.coins)
		}
	}
}
