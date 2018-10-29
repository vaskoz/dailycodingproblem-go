package day68

import "testing"

var testcases = []struct {
	bishops  []BoardPosition
	m        int
	expected int
}{
	{[]BoardPosition{{0, 0}, {1, 2}, {2, 2}, {4, 0}}, 5, 2},
}

func TestCountAttackingBishopPairs(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountAttackingBishopPairs(tc.bishops, tc.m); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkCountAttackingBishopPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountAttackingBishopPairs(tc.bishops, tc.m)
		}
	}
}
