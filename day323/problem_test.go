package day323

import (
	"math/rand"
	"testing"
)

// nolint
var testcases = []struct {
	unordered    []int
	approxMedian int
}{
	{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 7},
	{[]int{10, 10, 10, 10, 10, 1, 1, 1, 1, 1}, 5},
}

func TestApproximateMedian(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		r = rand.New(rand.NewSource(42)) // fixed number for test determinism.
		if median := ApproximateMedian(tc.unordered); median != tc.approxMedian {
			t.Errorf("Expected %v, got %v", tc.approxMedian, median)
		}
	}
}

func BenchmarkApproximateMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ApproximateMedian(tc.unordered)
		}
	}
}
