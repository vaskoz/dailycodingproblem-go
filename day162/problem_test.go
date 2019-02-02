package day162

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	words, prefixes []string
}{
	{[]string{"dog", "cat", "apple", "apricot", "fish"}, []string{"d", "c", "app", "apr", "f"}},
}

func TestShortestUniquePrefix(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ShortestUniquePrefix(tc.words); !reflect.DeepEqual(result, tc.prefixes) {
			t.Errorf("Expected %v got %v", tc.prefixes, result)
		}
	}
}

func BenchmarkShortestUniquePrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ShortestUniquePrefix(tc.words)
		}
	}
}
