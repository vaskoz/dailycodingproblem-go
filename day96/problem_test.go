package day96

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    []int
	expected [][]int
}{
	{[]int{1, 2, 3},
		[][]int{
			{1, 2, 3},
			{2, 1, 3},
			{3, 1, 2},
			{1, 3, 2},
			{2, 3, 1},
			{3, 2, 1}},
	},
}

func TestHeapsAlgorithm(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		results := HeapsAlgorithm(tc.input)
		pos := 0
		for result := range results {
			if !reflect.DeepEqual(result, tc.expected[pos]) {
				t.Errorf("Expected %v got %v", tc.expected[pos], result)
			}
			pos++
		}
	}
}

func BenchmarkHeapsAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			HeapsAlgorithm(tc.input)
		}
	}
}
