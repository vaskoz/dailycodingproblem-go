package day69

import "testing"

var testcases = []struct {
	input    []int
	expected int
}{
	{[]int{-10, -10, 5, 2}, 500},
	{[]int{2, 5, -10, -10, 5, 2}, 500},
	{[]int{2, 5, 10, 10, 6, 2}, 600},
	{[]int{2, 5, -10, -20, 6, 2}, 1200},
}

func TestLargestProductOf3(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LargestProductOf3(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkLargestProductOf3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestProductOf3(tc.input)
		}
	}
}
