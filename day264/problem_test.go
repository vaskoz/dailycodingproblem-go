package day264

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	c        string
	k        int
	expected string
}{
	{"01", 3, "00010111"},
	{"abcd", 2, "aabacadbbcbdccdd"},
}

func TestDeBruijnSequence(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := DeBruijnSequence(tc.c, tc.k); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkDeBruijnSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DeBruijnSequence(tc.c, tc.k)
		}
	}
}
