package day380

import "testing"

// nolint
var testcases = []struct {
	dividend, divisor   int
	quotient, remainder int
}{
	{10, 3, 3, 1},
	{9, 3, 3, 0},
	{-9, 3, -3, 0},
	{9, -3, -3, 0},
	{-9, -3, 3, 0},
}

func TestDivisionBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if q, r := DivisionBrute(tc.dividend, tc.divisor); q != tc.quotient || r != tc.remainder {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.quotient, tc.remainder, q, r)
		}
	}
}

func TestDivisionBruteDivideByZero(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected a panic when divisor is zero")
		}
	}()
	DivisionBrute(1, 0)
}

func BenchmarkDivisionBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DivisionBrute(tc.dividend, tc.divisor)
		}
	}
}
