package day316

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	ways, denoms []int
}{
	{[]int{1, 0, 1, 1, 2}, []int{2, 3, 4}},
}

func TestCoinDenominationsNeeded(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := CoinDenominationsNeeded(tc.ways); !reflect.DeepEqual(res, tc.denoms) {
			t.Errorf("Expected %v, got %v", tc.denoms, res)
		}
	}
}

func BenchmarkCoinDenominationsNeeded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CoinDenominationsNeeded(tc.ways)
		}
	}
}
