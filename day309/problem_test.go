package day309

import "testing"

// nolint
var testcases = []struct {
	seating  []int
	minMoves int
}{
	{[]int{0, 1, 1, 0, 1, 0, 0, 0, 1}, 5},
}

func TestMinRedistributeNoGapsBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if minMoves := MinRedistributeNoGapsBrute(tc.seating); minMoves != tc.minMoves {
			t.Errorf("Expected %v, got %v", tc.minMoves, minMoves)
		}
	}
}

func BenchmarkMinRedistributeNoGapsBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinRedistributeNoGapsBrute(tc.seating)
		}
	}
}
