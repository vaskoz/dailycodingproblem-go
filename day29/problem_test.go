package day29

import "testing"

var testcases = []struct {
	orig, encoded string
}{
	{"AAAABBBCCDAA", "4A3B2C1D2A"},
	{"AAAAAAAAAAA", "11A"},
}

func TestRunLength(t *testing.T) {
	for _, tc := range testcases {
		var enc string
		if enc = RunLengthEncoding(tc.orig); enc != tc.encoded {
			t.Errorf("Expected %v but got %v", tc.encoded, enc)
		}
		if dec := RunLengthDecoding(enc); dec != tc.orig {
			t.Errorf("Expected %v but got %v", tc.orig, dec)
		}
	}
}

func BenchmarkRunLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RunLengthDecoding(RunLengthEncoding(tc.orig))
		}
	}
}
