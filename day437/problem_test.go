package day437

import "testing"

// nolint
var testcases = []struct {
	str      string
	letters  []rune
	expected string
}{
	{"figehaeci", []rune("aei"), "aeci"},
	{"figehaeci", []rune("xei"), ""},
	{"this is a test string", []rune("tist"), "t stri"},
	{"a", []rune("ab"), ""},
}

func TestShortestSubstringContains(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := ShortestSubstringContains(tc.str, tc.letters); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkShortestSubstringContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ShortestSubstringContains(tc.str, tc.letters)
		}
	}
}
