package day184

import "testing"

var testcases = []struct {
	nums     []int
	expected int
}{
	{[]int{42, 56, 14}, 14},
	{[]int{14}, 14},
	{nil, 0},
	{[]int{}, 0},
}

func TestGcdSlice(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := GcdSlice(tc.nums); result != tc.expected {
			t.Errorf("Expected %v, got %v for input %v", tc.expected, result, tc.nums)
		}
	}
}

func BenchmarkGcdSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			GcdSlice(tc.nums)
		}
	}
}
