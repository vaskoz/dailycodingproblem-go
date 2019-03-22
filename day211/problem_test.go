package day211

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	str, pattern string
	expected     []int
}{
	{"abracadabra", "abr", []int{0, 7}},
}

func TestStartingIndicies(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := StartingIndicies(tc.str, tc.pattern); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkStartingIndicies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			StartingIndicies(tc.str, tc.pattern)
		}
	}
}
