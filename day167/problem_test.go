package day167

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	words    []string
	expected []Pair
}{
	{[]string{"code", "edoc", "da", "d"}, []Pair{{0, 1}, {1, 0}, {2, 3}}},
}

func TestPairsPalindromes(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := PairsPalindromes(tc.words); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkPairsPalindromes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PairsPalindromes(tc.words)
		}
	}
}
