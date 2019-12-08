package day383

import "testing"

// nolint
var testcases = []struct {
	s        string
	lst      []string
	expected string
}{
	{"abcdefg", []string{"bc", "ef"}, "a<b>bc</b>d<b>ef</b>g"},
	{"abcdefg", []string{"bcd", "def"}, "a<b>bcdef</b>g"},
	{"abcdefg", []string{"def", "bcd"}, "a<b>bcdef</b>g"},
	{"abcdefg", []string{"def", "bcd", "defg"}, "a<b>bcdefg</b>"},
	{"abcdefg", []string{"abcd", "bc"}, "<b>abcd</b>efg"},
}

func TestEmbolden(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := Embolden(tc.s, tc.lst); res != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkEmbolden(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Embolden(tc.s, tc.lst)
		}
	}
}
