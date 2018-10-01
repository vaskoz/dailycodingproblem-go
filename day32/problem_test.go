package day32

import "testing"

var testcases = []struct {
	rates             [][]float64
	arbitragePossible bool
}{
	{[][]float64{
		{1, 0.741, 0.657, 1.061, 1.005},
		{1.349, 1, 0.888, 1.433, 1.366},
		{1.521, 1.126, 1, 1.614, 1.538},
		{0.942, 0.698, 0.619, 1, 0.953},
		{0.995, 0.732, 0.650, 1.049, 1},
	}, true},
	{[][]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}, false},
}

func TestArbitrage(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := Arbitrage(tc.rates); result != tc.arbitragePossible {
			t.Errorf("TC%d Expected %v got %v", tcid, tc.arbitragePossible, result)
		}
	}
}

func BenchmarkArbitrage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Arbitrage(tc.rates)
		}
	}
}
