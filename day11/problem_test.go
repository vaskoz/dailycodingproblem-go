package day11

import (
	"reflect"
	"sort"
	"testing"
)

var testcases = []struct {
	input    []string
	query    string
	expected []string
}{
	{[]string{"dog", "deer", "deal"}, "de", []string{"deal", "deer"}},
	{[]string{"dog", "deer", "deal"}, "a", nil},
	{[]string{"dog", "deer", "doggone"}, "dog", []string{"dog", "doggone"}},
	{[]string{"abc", "abcd", "adec", "bde"}, "bde", []string{"bde"}},
}

func TestTrie(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		trie := NewTrie(tc.input)
		result := trie.Match(tc.query)
		sort.Strings(result)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v by got %v for query %v", tc.expected, result, tc.query)
		}
	}
}

func BenchmarkTrie(b *testing.B) {
	tries := make([]*Trie, len(testcases))
	for i := range tries {
		tries[i] = NewTrie(testcases[i].input)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := i % len(tries)
		tries[index].Match(testcases[index].query)
	}
}
