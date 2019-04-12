package day233

import "testing"

var testcases = []struct {
	n, fib uint
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{12, 144},
	{28, 317811},
	{44, 701408733},
	{87, 679891637638612258},
}

func TestFastFibonnaci(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := FastFibonnaci(tc.n); result != tc.fib {
			t.Errorf("Given n=%d, expected %d, got %d", tc.n, tc.fib, result)
		}
	}
}

func BenchmarkFastFibonnaci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FastFibonnaci(tc.n)
		}
	}
}
