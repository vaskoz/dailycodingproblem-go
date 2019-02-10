package day172

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	s        string
	words    []string
	expected []int
}{
	{"dogcatcatcodecatdog", []string{"cat", "dog"}, []int{0, 13}},
	{"barfoobazbitbyte", []string{"dog", "cat"}, nil},
	{"aabaa", []string{"aa", "ab"}, []int{1}},
}

func TestConcatenatedSubstringIndicies(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ConcatenatedSubstringIndicies(tc.s, tc.words); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkConcatenatedSubstringIndicies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ConcatenatedSubstringIndicies(tc.s, tc.words)
		}
	}
}
