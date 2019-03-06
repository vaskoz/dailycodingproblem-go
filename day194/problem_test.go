package day194

import "testing"

var testcases = []struct {
	p, q          []int
	intersections int
}{
	{[]int{3, 1}, []int{1, 3}, 1},
	{[]int{0, 10, 7, 4, 12, 9, 2}, []int{3, 5, 11, 6, 1, 13, 8}, 11},
}

func TestCountIntersectionsBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountIntersectionsBrute(tc.p, tc.q); result != tc.intersections {
			t.Errorf("Expected %v, got %v", tc.intersections, result)
		}
	}
}

func BenchmarkCountIntersectionsBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountIntersectionsBrute(tc.p, tc.q)
		}
	}
}
