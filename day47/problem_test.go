package day47

import "testing"

var testcases = []struct {
	prices    []int
	maxProfit int
}{
	{[]int{9, 11, 8, 5, 7, 10}, 5},
	{[]int{}, 0},
	{nil, 0},
}

func TestMaxProfitBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if profit := MaxProfitBrute(tc.prices); tc.maxProfit != profit {
			t.Errorf("Expected profit %d but got %d", tc.maxProfit, profit)
		}
	}
}

func BenchmarkMaxProfitBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxProfitBrute(tc.prices)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if profit := MaxProfit(tc.prices); tc.maxProfit != profit {
			t.Errorf("Expected profit %d but got %d", tc.maxProfit, profit)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxProfit(tc.prices)
		}
	}
}
