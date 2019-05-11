package day260

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	hints    []rune
	expected []int
}{
	{
		[]rune{'+', '+', '-', '+'},
		[]int{1, 2, 3, 0, 4},
	},
}

func TestUnjumble(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Unjumble(tc.hints); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkUnjumble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Unjumble(tc.hints)
		}
	}
}
