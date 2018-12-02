package day102

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    []int
	k        int
	expected []int
}{
	{[]int{1, 2, 3, 4, 5}, 9, []int{2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 16, nil},
	{[]int{1, 2, 3, 4, 5}, 15, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
		49, []int{4, 5, 6, 7, 8, 9, 10}},
}

func TestContiguousSumBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ContiguousSumBrute(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkContiguousSumBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ContiguousSumBrute(tc.input, tc.k)
		}
	}
}

func TestContiguousSumFaster(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ContiguousSumFaster(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkContiguousSumFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ContiguousSumFaster(tc.input, tc.k)
		}
	}
}
