package day248

import "testing"

// nolint
var testcases = []struct {
	a, b, max int64
}{
	{10, 20, 20},
	{20, 1, 20},
	{30, 2, 30},
	{1, -2, 1},
}

func TestMaxIf(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if max := MaxIf(tc.a, tc.b); max != tc.max {
			t.Errorf("Expected %v, got %v", tc.max, max)
		}
	}
}

func BenchmarkMaxIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxIf(tc.a, tc.b)
		}
	}
}

func TestMaxSubtractAndShift(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if max := MaxSubtractAndShift(tc.a, tc.b); max != tc.max {
			t.Errorf("Expected %v, got %v", tc.max, max)
		}
	}
}

func BenchmarkMaxSubtractAndShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxSubtractAndShift(tc.a, tc.b)
		}
	}
}

func TestMaxXor(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if max := MaxXor(tc.a, tc.b); max != tc.max {
			t.Errorf("Expected %v, got %v", tc.max, max)
		}
	}
}

func BenchmarkMaxXor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxXor(tc.a, tc.b)
		}
	}
}
