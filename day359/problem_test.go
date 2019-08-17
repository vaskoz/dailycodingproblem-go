package day359

import "testing"

// nolint
var testcases = []struct {
	anagram   string
	sortedInt string
	err       error
}{
	{"niesevehrtfeev", "357", nil},
	{"fourzeroninesixseven", "04679", nil},
	{"nineextra", "", ErrNotPossible()},
}

func TestConvertAnagramToInt(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := ConvertAnagramToInt(tc.anagram); result != tc.sortedInt || err != tc.err {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.sortedInt, tc.err, result, err)
		}
	}
}

func BenchmarkConvertAnagramToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ConvertAnagramToInt(tc.anagram) // nolint
		}
	}
}
