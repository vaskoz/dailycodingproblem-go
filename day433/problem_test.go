package day433

import "testing"

// nolint
var testcases = []struct {
	n, next int
}{
	{6, 9},
	{156, 163},
}

func TestNextIntSameSetBits(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if next := NextIntSameSetBits(tc.n); next != tc.next {
			t.Errorf("Expected %v, got %v", tc.next, next)
		}
	}
}

func BenchmarkNextIntSameSetBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NextIntSameSetBits(tc.n)
		}
	}
}
