package day81

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	number   string
	mapping  DigitMappings
	expected []string
}{
	{"23", DigitMappings{'2': {"a", "b", "c"}, '3': {"d", "e", "f"}},
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestPossiblePhoneLetters(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := PossiblePhoneLetters(tc.number, tc.mapping); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkPossiblePhoneLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PossiblePhoneLetters(tc.number, tc.mapping)
		}
	}
}
