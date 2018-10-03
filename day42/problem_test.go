package day42

import (
	"reflect"
	"sort"
	"testing"
)

var testcases = []struct {
	input    []int
	k        int
	expected []int
}{
	{[]int{12, 1, 61, 5, 9, 2}, 24, []int{1, 2, 9, 12}},
	{[]int{12, 1, 61, 5, 9, 2}, 1000, nil},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		1000, nil},
}

func TestSubsetSumBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		result := SubsetSumBrute(tc.input, tc.k)
		sort.Ints(result)
		if !reflect.DeepEqual(tc.expected, result) {
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

func TestSubsetSumDP(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		result := SubsetSumDP(tc.input, tc.k)
		sort.Ints(result)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSubsetSumDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SubsetSumDP(tc.input, tc.k)
		}
	}
}
