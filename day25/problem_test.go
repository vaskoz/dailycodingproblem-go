package day25

import "testing"

var testcases = []struct {
	input, pattern string
	expected       bool
}{
	{"ray", "ra.", true},
	{"raymond", "ra.", false},
	{"chat", ".*at", true},
	{"chats", ".*at", false},
	{"chaaaa", "cha*", true},
	{"raymond", "ra.*", true},
	{"raymond", "ra.*d", true},
	{"chaaab", "cha*", false},
	{"chaaa", "cha*b", false},
	{"abcd", "efg", false},
	{"ab", "e", false},
}

func TestMatch(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Match(tc.input, tc.pattern); result != tc.expected {
			t.Errorf("Given (%v, %v), expected %v got %v", tc.input, tc.pattern, tc.expected, result)
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Match(tc.input, tc.pattern)
		}
	}
}
