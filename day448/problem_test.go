package day448

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input, expected []rune
}{
	{[]rune{'G', 'B', 'R', 'R', 'B', 'R', 'G'}, []rune{'R', 'R', 'R', 'G', 'G', 'B', 'B'}},
	{[]rune{'G', 'B', 'R', 'G'}, []rune{'R', 'G', 'G', 'B'}},
	{[]rune{'B', 'G', 'R'}, []rune{'R', 'G', 'B'}},
}

func TestSegregateValues(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		result := append(make([]rune, 0, len(tc.input)), tc.input...)
		SegregateValues(result)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkSegregateValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			result := append(make([]rune, 0, len(tc.input)), tc.input...)
			SegregateValues(result)
		}
	}
}
