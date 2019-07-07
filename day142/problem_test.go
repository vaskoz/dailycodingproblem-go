package day142

import "testing"

// nolint
var testcases = []struct {
	input    string
	balanced bool
}{
	{"(()*", true},
	{"(*)", true},
	{")*(", false},
	{"(((**", false},
	{"((***))", true},
	{"(***)))", true},
	{"**))**()))*", true},
	{"************(*)***********", true},
	{"***********************", true},
	{"****)))", true},
	{"****)))))", false},
	{"*(", false},
}

func TestWildcardParens(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if balanced := WildcardParens(tc.input); balanced != tc.balanced {
			t.Errorf("Given %v Expected %v got %v", tc.input, tc.balanced, balanced)
		}
	}
}

func BenchmarkWildcardParens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			WildcardParens(tc.input)
		}
	}
}
