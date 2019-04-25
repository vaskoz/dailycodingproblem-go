package day246

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input, expected []string
}{
	{
		[]string{"chair", "height", "racket", "touch", "tunic"},
		[]string{"chair", "racket", "touch", "height", "tunic"},
	},
	{[]string{"one"}, []string{"one"}},
	{
		[]string{"chair", "height", "touch", "tunic"},
		nil,
	},
	{[]string{"a", "aa", "ac", "ad", "ba", "ca", "da"}, nil},
}

func TestCircleWords(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CircleWords(tc.input); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkCircleWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CircleWords(tc.input)
		}
	}
}
