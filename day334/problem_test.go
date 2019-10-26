package day334

import "testing"

// nolint
var testcases = []struct {
	digits   []int
	expected bool
}{
	{[]int{5, 2, 7, 8}, true},
}

func TestTwentyFourGame(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := TwentyFourGame(tc.digits); result != tc.expected {
			t.Errorf("Given %v, expected %v, got %v", tc.digits, tc.expected, result)
		}
	}
}

func BenchmarkTwentyFourGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TwentyFourGame(tc.digits)
		}
	}
}
