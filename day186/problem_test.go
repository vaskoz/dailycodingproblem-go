package day186

import (
	"testing"
)

var testcases = []struct {
	nums []int
	diff int
}{
	{[]int{5, 10, 15, 20, 25}, 5},
}

func TestSubsetsSmallestDifference(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if s1, s2 := SubsetsSmallestDifference(tc.nums); sum(s1, s2) != tc.diff {
			t.Errorf("Expected a diff of %v, got (%v,%v) for a diff of %v", tc.diff, s1, s2, sum(s1, s2))
		}
	}
}

func BenchmarkSubsetsSmallestDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SubsetsSmallestDifference(tc.nums)
		}
	}
}
