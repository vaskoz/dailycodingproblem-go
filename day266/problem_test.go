package day266

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	dict           []string
	input          string
	validStepWords []string
}{
	{[]string{"appeal", "apearr", "appear"}, "apple", []string{"appeal"}},
}

func TestAllValidStepWords(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := AllValidStepWords(tc.dict, tc.input); !reflect.DeepEqual(result, tc.validStepWords) {
			t.Errorf("Expected %v, got %v", tc.validStepWords, result)
		}
	}
}

func BenchmarkAllValidStepWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AllValidStepWords(tc.dict, tc.input)
		}
	}
}
