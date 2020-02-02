package day430

import "testing"

// nolint
var testcases = []struct {
	parens   string
	expected string
}{
	{"(()", "(())"},
	{"))()(", "()()()()"},
	{"))(", "()()()"},
}

func TestBalanceParens(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := BalanceParens(tc.parens); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkBalanceParens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BalanceParens(tc.parens)
		}
	}
}
