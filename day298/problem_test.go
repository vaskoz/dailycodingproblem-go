package day298

import "testing"

// nolint
var testcases = []struct {
	apples   []int
	expected int
}{
	{[]int{2, 1, 2, 3, 3, 1, 3, 5}, 4},
	{[]int{2, 1, 2, 2, 2, 1, 2, 1}, 8},
	{[]int{4, 3, 2, 1}, 2},
}

func TestLongestAppleRun(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LongestAppleRun(tc.apples); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkLongestAppleRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestAppleRun(tc.apples)
		}
	}
}
