package day61

import "testing"

var testcases = []struct {
	x, y, expected int
}{
	{2, 10, 1024},
}

func TestNaiveIntegerExponentiation(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := NaiveIntegerExponentiation(tc.x, tc.y); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func TestNaiveIntegerExponentiationPanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if x := recover(); x == nil {
			t.Errorf("Expected a panic for negative exponent")
		}
	}()
	NaiveIntegerExponentiation(1, -1)
}

func BenchmarkNaiveIntegerExponentiation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NaiveIntegerExponentiation(tc.x, tc.y)
		}
	}
}
