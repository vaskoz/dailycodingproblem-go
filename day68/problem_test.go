package day68

import "testing"

// nolint
var testcases = []struct {
	bishops  []BoardPosition
	expected int
}{
	{[]BoardPosition{{0, 0}, {1, 2}, {2, 2}, {4, 0}}, 2},
	{[]BoardPosition{{0, 0}, {1, 4}, {2, 2}, {0, 1}, {4, 0}, {2, 3}, {4, 4}}, 6},
}

func TestCountAttackingBishopPairs(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := CountAttackingBishopPairs(tc.bishops); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkCountAttackingBishopPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountAttackingBishopPairs(tc.bishops)
		}
	}
}
