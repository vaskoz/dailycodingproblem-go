package day206

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input                []interface{}
	permutationPositions []int
	expected             []interface{}
}{
	{[]interface{}{"a", "b", "c"}, []int{2, 1, 0}, []interface{}{"c", "b", "a"}},
}

func TestReorderPermutation(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ReorderPermutation(tc.input, tc.permutationPositions); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkReorderPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ReorderPermutation(tc.input, tc.permutationPositions)
		}
	}
}
