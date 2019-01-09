package day121

import "testing"

var testcases = []struct {
	input    string
	k        int
	expected bool
}{
	{"waterrfetawx", 2, true},
}

func TestIsKPalindromeBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsKPalindromeBrute(tc.input, tc.k); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkIsKPalindromeBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsKPalindromeBrute(tc.input, tc.k)
		}
	}
}
