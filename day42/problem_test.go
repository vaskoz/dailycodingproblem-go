package day42

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    []int
	k        int
	expected []int
}{
	{[]int{12, 1, 61, 5, 9, 2}, 24, []int{12, 1, 9, 2}},
	{[]int{12, 1, 61, 5, 9, 2}, 1000, nil},
}

func TestSubsetSumBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SubsetSumBrute(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSubsetSumBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SubsetSumBrute(tc.input, tc.k)
		}
	}
}
