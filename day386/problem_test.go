package day386

import "testing"

// nolint
var testcases = []struct {
	input, expected string
}{
	{"tweet", "tteew"},
	{"happiness", "ssppnihea"},
}

func TestSortStringByCharFreq(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := SortStringByCharFreq(tc.input); res != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkSortStringByCharFreq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SortStringByCharFreq(tc.input)
		}
	}
}
