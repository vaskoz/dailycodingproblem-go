package day342

import (
	"fmt"
	"testing"
)

// nolint
var testcases = []struct {
	lst      []interface{}
	reducer  func(a, b interface{}) interface{}
	initial  interface{}
	expected interface{}
}{
	{
		[]interface{}{1, 2, 3, 4, 5, 6},
		func(a, b interface{}) interface{} {
			return a.(int) + b.(int)
		},
		0,
		21,
	},
	{
		[]interface{}{"hello", "world", "goodbye"},
		func(a, b interface{}) interface{} {
			return fmt.Sprintf("%s %s", a, b)
		},
		"begin",
		"begin hello world goodbye",
	},
}

func TestReduce(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Reduce(tc.lst, tc.reducer, tc.initial); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Reduce(tc.lst, tc.reducer, tc.initial)
		}
	}
}
