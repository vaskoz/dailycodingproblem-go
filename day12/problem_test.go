package day12

import "testing"

var testcases = []struct {
	steps    int
	strides  []int
	expected int
}{
	{4, []int{1, 2}, 5},
	{5, []int{1, 2}, 8},
	{5, []int{1, 2}, 8},
	{4, []int{1, 2, 3}, 7},
	{3, []int{1, 2, 3}, 4},
	{100, []int{1, 100}, 2},
	{100, []int{1, 50, 100}, 54},
	{100, []int{50, 1, 100}, 54},
	{4, []int{2, 1}, 5},
	{4, []int{2}, 1},
	{1, []int{3, 2}, 0},
}

func TestUniqueClimbs(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := UniqueClimbs(tc.steps, tc.strides); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkUniqueClimbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			UniqueClimbs(tc.steps, tc.strides)
		}
	}
}

func TestUniqueClimbsDS(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := UniqueClimbsDS(tc.steps, tc.strides); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkUniqueClimbsDS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			UniqueClimbsDS(tc.steps, tc.strides)
		}
	}
}
