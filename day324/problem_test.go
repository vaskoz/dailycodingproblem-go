package day324

import "testing"

// nolint
var testcases = []struct {
	mice, holes []int
	expected    int
}{
	{[]int{1, 4, 9, 15}, []int{10, -5, 0, 16}, 6},
}

func TestMostMouseSteps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MostMouseSteps(tc.mice, tc.holes); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestMostMouseStepsUnmatched(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic from unmatched inputs")
		}
	}()
	MostMouseSteps([]int{10}, []int{10, 20})
}

func BenchmarkMostMouseSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MostMouseSteps(tc.mice, tc.holes)
		}
	}
}
