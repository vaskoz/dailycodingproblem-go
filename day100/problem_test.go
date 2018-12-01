package day100

import "testing"

var testcases = []struct {
	path         []Point
	minimumSteps int
}{
	{[]Point{{0, 0}, {1, 1}, {1, 2}}, 2},
	{[]Point{}, 0},
	{nil, 0},
	{[]Point{{100, 100}}, 0},
	{[]Point{{0, 0}, {-10, -10}, {20, 20}}, 40},
	{[]Point{{0, 0}, {-10, -10}, {20, 40}}, 60},
}

func TestMinimumStepsPath(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if steps := MinimumStepsPath(tc.path); steps != tc.minimumSteps {
			t.Errorf("Expected %v got %v", tc.minimumSteps, steps)
		}
	}
}

func BenchmarkMinimumStepsPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumStepsPath(tc.path)
		}
	}
}
