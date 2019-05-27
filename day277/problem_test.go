package day277

import "testing"

// nolint
var testcases = []struct {
	bits  string
	valid bool
}{
	{"111000101000001010101100", true},
	{"0111111100000000", true},
	{"110000001000000001010101", true},
	{"11110111101111111011100010000111", true},
	{"11110111", false},
	{"10000000", false},
	{"11000000", false},
	{"11100000", false},
	{"1110000001010101", false},
	{"1100000001010101", false},
}

func TestIsValidUTF8(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsValidUTF8(tc.bits); result != tc.valid {
			t.Errorf("Given %s, expected %v, got %v", tc.bits, tc.valid, result)
		}
	}
}

func TestIsValidUTF8BadInput(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected a panic with bad input")
		}
	}()
	IsValidUTF8("10111")
}

func BenchmarkIsValidUTF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsValidUTF8(tc.bits)
		}
	}
}
