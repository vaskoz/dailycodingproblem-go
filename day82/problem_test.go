package day82

import (
	"reflect"
	"strings"
	"testing"
)

var testcases = []struct {
	fileData string
	n        int
	reads    int
	expected []string
}{
	{"Hello world", 7, 3, []string{"Hello w", "orld", ""}},
	{"Hi", 2, 6, []string{"Hi", "", "", "", "", ""}},
}

func TestReadN(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		r := strings.NewReader(tc.fileData)
		results := make([]string, 0, tc.reads)
		for i := 0; i < tc.reads; i++ {
			results = append(results, ReadN(r, tc.n))
		}
		if !reflect.DeepEqual(tc.expected, results) {
			t.Errorf("Expected %v got %v", tc.expected, results)
		}
	}
}

func BenchmarkReadN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			r := strings.NewReader(tc.fileData)
			for i := 0; i < tc.reads; i++ {
				ReadN(r, tc.n)
			}
		}
	}
}
