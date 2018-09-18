package day28

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	words    []string
	k        int
	expected []string
}{
	{[]string{"bird"},
		6,
		[]string{"bird  "}},
	{[]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
		16,
		[]string{"the  quick brown", "fox  jumps  over", "the   lazy   dog"}},
}

func TestJustify(t *testing.T) {
	for _, tc := range testcases {
		if result := Justify(tc.words, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkJustify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Justify(tc.words, tc.k)
		}
	}
}
