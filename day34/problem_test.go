package day34

import "testing"

var testcases = []struct {
	input, expected string
}{
	{"race", "ecarace"},
	{"google", "elgoogle"},
	{"add", "adda"},
	{"caoobac", "caboobac"},
	{"abc", "abcba"},
}

func TestMakePalindrome(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MakePalindrome(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkMakePalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MakePalindrome(tc.input)
		}
	}
}
