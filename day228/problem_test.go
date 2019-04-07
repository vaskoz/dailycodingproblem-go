package day228

import "testing"

var testcases = []struct {
	nums     []int
	expected int
}{
	{[]int{10, 7, 76, 415}, 77641510},
	{[]int{7, 76, 789}, 789776},
	{[]int{789, 7, 76}, 789776},
	{[]int{789, 76, 7}, 789776},
	{[]int{791, 7}, 7917},
}

func TestLargestPossibleInteger(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LargestPossibleInteger(tc.nums); result != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, result)
		}
	}
}

func BenchmarkLargestPossibleInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestPossibleInteger(tc.nums)
		}
	}
}
