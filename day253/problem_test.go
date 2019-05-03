package day253

import (
	"bytes"
	"testing"
)

// nolint
var testcases = []struct {
	letters  string
	k        int
	expected string
}{
	{"thisisazigzag", 4, "t     a     g\n h   s z   a \n  i i   i z  \n   s     g   \n"},
}

func TestPrintZigZag(t *testing.T) {
	// don't run in parallel due to global var
	for _, tc := range testcases {
		buff := new(bytes.Buffer)
		out = buff
		PrintZigZag(tc.letters, tc.k)
		if result := buff.String(); result != tc.expected {
			t.Errorf("Expected \n%s, got \n%s", tc.expected, result)
		}
	}
}

func BenchmarkPrintZigZag(b *testing.B) {
	// don't run in parallel due to global var
	buff := new(bytes.Buffer)
	out = buff
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PrintZigZag(tc.letters, tc.k)
		}
	}
}
