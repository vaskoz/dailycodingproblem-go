package day401

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input    []interface{}
	perm     []int
	expected []interface{}
}{
	{
		[]interface{}{"a", "b", "c"},
		[]int{2, 1, 0},
		[]interface{}{"c", "b", "a"},
	},
}

func TestApplyPermutation(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		copied := append([]interface{}{}, tc.input...)
		if ApplyPermutation(copied, tc.perm); !reflect.DeepEqual(copied, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, copied)
		}
	}
}

func TestApplyPermutationBadInput(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic")
		}
	}()
	ApplyPermutation([]interface{}{"a", "b", "c"}, []int{2, 1, 0, 3})
}

func BenchmarkApplyPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := append([]interface{}{}, tc.input...)
			ApplyPermutation(copied, tc.perm)
		}
	}
}
