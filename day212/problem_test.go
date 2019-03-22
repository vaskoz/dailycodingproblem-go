package day212

import "testing"

var testcases = []struct {
	pos  int
	cell string
}{
	{1, "A"},
	{27, "AA"},
	{52, "AZ"},
	{53, "BA"},
	{78, "BZ"},
	{79, "CA"},
	{677, "ZA"},
	{702, "ZZ"},
	{703, "AAA"},
	{0, ""},
	{-10, ""},
}

func TestSpreadsheetColumnName(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SpreadsheetColumnName(tc.pos); result != tc.cell {
			t.Errorf("Expected %v, got %v", tc.cell, result)
		}
	}
}

func BenchmarkSpreadsheetColumnName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SpreadsheetColumnName(tc.pos)
		}
	}
}
