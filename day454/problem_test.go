package day454

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	seq []int
	lis []int
}{
	{[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}, []int{0, 2, 6, 9, 11, 15}},
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := LongestIncreasingSubsequence(tc.seq); !reflect.DeepEqual(result, tc.lis) {
			t.Errorf("Expected %v, got %v", tc.lis, result)
		}
	}
}

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestIncreasingSubsequence(tc.seq)
		}
	}
}
