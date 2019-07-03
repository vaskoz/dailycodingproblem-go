package day314

import (
	"testing"
)

// nolint
var testcases = []struct {
	listeners, towers []int
	minRange          int
}{
	{[]int{1, 5, 11, 20}, []int{4, 8, 15}, 5},
	{[]int{20, 5, 1, 11}, []int{15, 8, 4}, 5},
	{[]int{1, 5, 11, 20}, []int{21}, 20},
	{[]int{1, 5, 11, 20}, []int{0}, 20},
}

func TestMinimumTowerRange(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinimumTowerRange(tc.listeners, tc.towers); result != tc.minRange {
			t.Errorf("Expected %v, got %v", tc.minRange, result)
		}
	}
}

func BenchmarkMinimumTowerRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumTowerRange(tc.listeners, tc.towers)
		}
	}
}
