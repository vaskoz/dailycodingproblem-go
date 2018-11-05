package day75

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input, expected []int
}{
	{[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
		[]int{0, 2, 6, 9, 11, 15}},
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LongestIncreasingSubsequence(tc.input); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestIncreasingSubsequence(tc.input)
		}
	}
}
