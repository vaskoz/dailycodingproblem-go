package day408

import "testing"

// nolint
var testcases = []struct {
	prices    []int
	k         int
	maxProfit int
}{
	{[]int{5, 2, 4, 0, 1}, 2, 3},
	{[]int{12, 14, 17, 10, 14, 13, 12, 15}, 3, 12},
}

func TestMaxProfit(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if maxProfit := MaxProfit(tc.prices, tc.k); maxProfit != tc.maxProfit {
			t.Errorf("Expected %v got %v", tc.maxProfit, maxProfit)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxProfit(tc.prices, tc.k)
		}
	}
}
