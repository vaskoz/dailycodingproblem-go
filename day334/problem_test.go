package day334

import "testing"

// nolint
var testcases = []struct {
	digits   []int
	expected bool
}{
	{[]int{5, 2, 7, 8}, true},
	{[]int{100, 100, 100, 100, 100}, false},
	{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2}, true},
	{[]int{5760000, 10, 30, 40, 20}, true},
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
