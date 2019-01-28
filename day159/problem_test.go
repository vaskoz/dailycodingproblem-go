package day159

import "testing"

var testcases = []struct {
	s           string
	expected    rune
	expectedErr error
}{
	{"acbbac", 'b', nil},
	{"abcdef", 0, ErrNoRecurringRune()},
}

func TestFirstRecurringRune(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if r, e := FirstRecurringRune(tc.s); r != tc.expected || e != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, r, e)
		}
	}
}

func BenchmarkFirstRecurringRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FirstRecurringRune(tc.s)
		}
	}
}
