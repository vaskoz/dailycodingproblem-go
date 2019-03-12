package day202

import "testing"

var testcases = []struct {
	num      int
	expected bool
}{
	{888, true},
	{121, true},
	{678, false},
}

func TestIsIntegerPalindrome(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsIntegerPalindrome(tc.num); result != tc.expected {
			t.Errorf("For input %v, expected %v, got %v", tc.num, tc.expected, result)
		}
	}
}

func BenchmarkIsIntegerPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsIntegerPalindrome(tc.num)
		}
	}
}
