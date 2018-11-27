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

func TestHeapsAlgorithmRecursive(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		results := HeapsAlgorithmRecursive(tc.input)
		pos := 0
		for result := range results {
			if !reflect.DeepEqual(result, tc.expected[pos]) {
				t.Errorf("Expected %v got %v", tc.expected[pos], result)
			}
			pos++
		}
	}
}

func BenchmarkHeapsAlgorithmRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			HeapsAlgorithmRecursive(tc.input)
		}
	}
}

func TestHeapsAlgorithmIterative(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		results := HeapsAlgorithmIterative(tc.input)
		pos := 0
		for result := range results {
			if !reflect.DeepEqual(result, tc.expected[pos]) {
				t.Errorf("Expected %v got %v", tc.expected[pos], result)
			}
			pos++
		}
	}
}

func BenchmarkHeapsAlgorithmIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			HeapsAlgorithmIterative(tc.input)
		}
	}
}
