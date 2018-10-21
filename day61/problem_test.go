package day61

import "testing"

var testcases = []struct {
	x, y, expected uint
}{
	{2, 10, 1024},
	{2, 1, 2},
	{2, 0, 1},
	{2, 62, 4611686018427387904},
}

func TestNaiveIntegerExponentiation(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := NaiveIntegerExponentiation(tc.x, tc.y); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkNaiveIntegerExponentiation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NaiveIntegerExponentiation(tc.x, tc.y)
		}
	}
}

func TestIntegerExponentiationBySquaring(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IntegerExponentiationBySquaring(tc.x, tc.y); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkIntegerExponentiationBySquaring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IntegerExponentiationBySquaring(tc.x, tc.y)
		}
	}
}
