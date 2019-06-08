package day290

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input    []rune
	expected []rune
}{
	{[]rune{'R', 'G', 'B', 'G', 'B'}, []rune{'R'}},
}

func TestMinimizeQux(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinimizeQux(tc.input); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkMinimizeQux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimizeQux(tc.input)
		}
	}
}
