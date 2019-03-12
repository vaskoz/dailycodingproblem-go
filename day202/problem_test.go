package day202

import "testing"

var testcases = []struct {
	num      int
	expected bool
}{
	{888, true},
	{121, true},
	{678, false},
	{123124234323, false},
	{123454321, true},
	{23455432, true},
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

func TestIsIntegerPalindromeFaster(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsIntegerPalindromeFaster(tc.num); result != tc.expected {
			t.Errorf("For input %v, expected %v, got %v", tc.num, tc.expected, result)
		}
	}
}

func BenchmarkIsIntegerPalindromeFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsIntegerPalindromeFaster(tc.num)
		}
	}
}
