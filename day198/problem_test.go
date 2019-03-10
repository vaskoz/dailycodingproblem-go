package day198

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	nums, subset []int
}{
	{[]int{3, 5, 10, 20, 21}, []int{5, 10, 20}},
	{[]int{1, 3, 6, 24}, []int{1, 3, 6, 24}},
}

func TestLargestSubsetPairs(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		input := append([]int{}, tc.nums...)
		if result := LargestSubsetPairs(input); !reflect.DeepEqual(result, tc.subset) {
			t.Errorf("Expected %v, got %v", tc.subset, result)
		}
	}
}

func BenchmarkLargestSubsetPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			input := append([]int{}, tc.nums...)
			LargestSubsetPairs(input)
		}
	}
}
