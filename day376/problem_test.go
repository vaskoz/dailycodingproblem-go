package day376

import "testing"

// nolint
var testcases = []struct {
	coins    []Pos
	me       Pos
	expected Pos
}{
	{[]Pos{{0, 4}, {1, 0}, {2, 0}, {3, 2}}, Pos{0, 2}, Pos{0, 4}},
}

func TestClosestCoinToMe(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := ClosestCoinToMe(tc.coins, tc.me); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkClosestCoinToMe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ClosestCoinToMe(tc.coins, tc.me)
		}
	}
}
