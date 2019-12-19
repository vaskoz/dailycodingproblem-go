package day395

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	words    []string
	expected [][]string
}{
	{
		[]string{"eat", "ate", "apt", "pat", "tea", "now"},
		[][]string{
			{"eat", "ate", "tea"},
			{"apt", "pat"},
			{"now"},
		},
	},
}

func TestGroupByAnagram(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := GroupByAnagram(tc.words); !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkGroupByAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			GroupByAnagram(tc.words)
		}
	}
}
