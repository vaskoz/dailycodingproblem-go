package day109

import "testing"

var testcases = []struct {
	ui       uint8
	expected uint8
}{
	{170, 85},
	{226, 209},
}

func TestSwapEvenOddBits(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SwapEvenOddBits(tc.ui); result != tc.expected {
			t.Errorf("Given %v expected %v got %v", tc.ui, tc.expected, result)
		}
	}
}

func BenchmarkSwapEvenOddBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SwapEvenOddBits(tc.ui)
		}
	}
}
