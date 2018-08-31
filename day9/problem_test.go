package day9

import "testing"

var testcases = []struct {
	input    []int
	expected int
}{
	{[]int{2, 4, 6, 2, 5}, 13},
	{[]int{5, 1, 1, 5}, 10},
	{[]int{10, 1}, 10},
}

func TestMaximumNonAdjacentSum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaximumNonAdjacentSum(tc.input); result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func BenchmarkMaximumNonAdjacentSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaximumNonAdjacentSum(tc.input)
		}
	}
}
