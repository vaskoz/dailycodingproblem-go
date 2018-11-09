package day79

import "testing"

var testcases = []struct {
	input    []int
	expected bool
}{
	{[]int{10, 5, 7}, true},
	{[]int{10, 5, 1}, false},
	{[]int{6, 10, 5, 6}, false},
	{[]int{6, 10, 5}, true},
}

func TestCouldBeNonDecreasingByOneChange(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := CouldBeNonDecreasingByOneChange(tc.input); result != tc.expected {
			t.Errorf("TC%d Expected %v got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkCouldBeNonDecreasingByOneChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CouldBeNonDecreasingByOneChange(tc.input)
		}
	}
}
