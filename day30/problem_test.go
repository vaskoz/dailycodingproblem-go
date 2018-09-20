package day30

import "testing"

var testcases = []struct {
	elevations []int
	expected   int
}{
	{[]int{2, 1}, 0},
	{[]int{2, 1, 2}, 1},
	{[]int{3, 0, 1, 3, 0, 5}, 8},
	{[]int{1, 2, 3, 0, 5}, 3},
}

func TestTrappedWater(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := TrappedWater(tc.elevations); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkTrappedWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TrappedWater(tc.elevations)
		}
	}
}
