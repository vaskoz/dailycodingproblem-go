package day457

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	W, S     string
	expected []int
}{
	{"ab", "abxaba", []int{0, 3, 4}},
	{"ba", "abxaba", []int{0, 3, 4}},
	{"aabx", "abxaba", []int{0, 2}},
}

func TestStartIndicesAnagrams(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := StartIndicesAnagrams(tc.W, tc.S); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}
func BenchmarkStartIndicesAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			StartIndicesAnagrams(tc.W, tc.S)
		}
	}
}
