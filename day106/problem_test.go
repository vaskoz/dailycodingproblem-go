package day106

import "testing"

var testcases = []struct {
	hops     []int
	expected bool
}{
	{[]int{2, 0, 1, 0}, true},
	{[]int{1, 1, 0, 1}, false},
}

func TestCanHopToEnd(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CanHopToEnd(tc.hops); result != tc.expected {
			t.Errorf("For %v expected %v got %v", tc.hops, tc.expected, result)
		}
	}
}

func BenchmarkCanHopToEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CanHopToEnd(tc.hops)
		}
	}
}
