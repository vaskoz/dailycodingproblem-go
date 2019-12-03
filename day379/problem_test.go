package day379

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input  string
	subseq map[string]struct{}
}{
	{"xyz",
		map[string]struct{}{
			"x":   {},
			"y":   {},
			"z":   {},
			"xy":  {},
			"xz":  {},
			"yz":  {},
			"xyz": {},
		}},
}

func TestAllPossibleSubsequences(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := AllPossibleSubsequences(tc.input); !reflect.DeepEqual(result, tc.subseq) {
			t.Errorf("Expected %v, got %v", tc.subseq, result)
		}
	}
}

func BenchmarkAllPossibleSubsequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AllPossibleSubsequences(tc.input)
		}
	}
}
