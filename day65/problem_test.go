package day65

import "testing"

var testcases = []struct {
	matrix   [][]int
	expected string
}{
	{
		[][]int{{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20}},
		"1 2 3 4 5 10 15 20 19 18 17 16 11 6 7 8 9 14 13 12",
	},
	{
		[][]int{{1, 2},
			{4, 3}},
		"1 2 3 4",
	},
	{
		[][]int{{}},
		"",
	},
	{
		[][]int{},
		"",
	},
}

func TestClockwiseSpiral(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ClockwiseSpiral(tc.matrix); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkClockwiseSpiral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ClockwiseSpiral(tc.matrix)
		}
	}
}
