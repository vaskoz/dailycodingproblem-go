package day315

import "testing"

// nolint
var testcases = []struct {
	mat      [][]int
	expected bool
}{
	{[][]int{
		{1, 2, 3, 4, 8},
		{5, 1, 2, 3, 4},
		{4, 5, 1, 2, 3},
		{7, 4, 5, 1, 2},
	},
		true,
	},
	{[][]int{
		{1, 2, 3, 4, 8},
		{5, 1, 2, 3, 4},
		{4, 5, 1, 2, 3},
		{7, 4, 5, 9, 2},
	},
		false,
	},
	{[][]int{
		{1, 2, 3, 4, 8},
		{5, 1, 2, 3, 4},
		{4, 5, 1, 2, 3},
		{7, 4, 5, 1, 9},
	},
		false,
	},
}

func TestIsToeplitzMatrix(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := IsToeplitzMatrix(tc.mat); result != tc.expected {
			t.Errorf("Expected %v, got %v for TCID %d", tc.expected, result, tcid)
		}
	}
}

func BenchmarkIsToeplitzMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsToeplitzMatrix(tc.mat)
		}
	}
}
