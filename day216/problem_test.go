package day216

import "testing"

var testcases = []struct {
	roman    string
	expected int
}{
	{"IV", 4},
	{"XL", 40},
	{"XIV", 14},
	{"MMXIX", 2019},
	{"XVI", 16},
	{"XMMMM", 3990},
}

func TestConvertRomanNumeralToInteger(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ConvertRomanNumeralToInteger(tc.roman); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkConvertRomanNumeralToInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ConvertRomanNumeralToInteger(tc.roman)
		}
	}
}
