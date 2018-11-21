package day90

import "testing"

var testcases = []struct {
	n            int
	l            []int
	distribution float64
}{
	{5, []int{0, 1, 2, 2}, 0.5},
	{5, []int{0, 1, 0, 1}, 0.333},
	{5, []int{0, 0, 0}, 0.25},
	{5, []int{0, 1, 2, 3}, 1.0},
}

func TestRandMissingNumbers(t *testing.T) {
	//t.Parallel()//don't run in parallel
	for _, tc := range testcases {
		results := make(map[int]int)
		const iterations = 1000
		for i := 0; i < iterations; i++ {
			results[RandMissingNumbers(tc.n, tc.l)]++
		}
		for _, v := range results {
			delta := (float64(v) / float64(iterations)) - tc.distribution
			if delta < 0 {
				delta = -delta
			}
			if delta > 0.1 {
				t.Errorf("Exceeded 0.1 tolerance for expected %v distribution", tc.distribution)
			}
		}
	}
}

func TestRandMissingNoNumbers(t *testing.T) {
	//t.Parallel()//don't run in parallel
	if result := RandMissingNumbers(5, []int{0, 1, 2, 3, 4}); result != -1 {
		t.Errorf("Expected -1 for a list that includes all possible numbers")
	}
}

func BenchmarkRandMissingNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RandMissingNumbers(tc.n, tc.l)
		}
	}
}
