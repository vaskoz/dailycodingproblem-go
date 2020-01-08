package day415

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	prices    []int
	fee       int
	maxProfit int
}{
	{[]int{1, 3, 2, 8, 4, 10}, 2, 9},
}

func TestMaxProfitWithFee(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := MaxProfitWithFee(tc.prices, tc.fee); !reflect.DeepEqual(tc.maxProfit, result) {
			t.Errorf("Expected %v, got %v", tc.maxProfit, result)
		}
	}
}

func BenchmarkMaxProfitWithFee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxProfitWithFee(tc.prices, tc.fee)
		}
	}
}
