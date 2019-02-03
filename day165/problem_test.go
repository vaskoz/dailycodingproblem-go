package day165

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	nums, expected []int
}{
	{[]int{3, 4, 9, 6, 1}, []int{1, 1, 2, 1, 0}},
}

func TestSmallerRightCount(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallerRightCount(tc.nums); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSmallerRightCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallerRightCount(tc.nums)
		}
	}
}
