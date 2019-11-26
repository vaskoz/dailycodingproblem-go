package day372

import "testing"

// nolint
var testcases = []struct {
	n         int
	numDigits int
}{
	{0, 1},
	{3, 1},
	{9, 1},
	{10, 2},
	{-10, 2},
	{-100, 3},
	{-99, 2},
	{99, 2},
}

func TestNumberOfDigits(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if numDigits := NumberOfDigits(tc.n); numDigits != tc.numDigits {
			t.Errorf("Given %v, expected %v, got %v", tc.n, tc.numDigits, numDigits)
		}
	}
}

func BenchmarkNumberOfDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumberOfDigits(tc.n)
		}
	}
}
