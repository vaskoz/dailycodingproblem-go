package day76

import "testing"

var testcases = []struct {
	input    [][]rune
	expected int
}{
	{[][]rune{
		[]rune("cba"),
		[]rune("daf"),
		[]rune("ghi"),
	}, 1},
	{[][]rune{
		[]rune("abcdef"),
	}, 0},
	{[][]rune{
		[]rune("zyx"),
		[]rune("wvu"),
		[]rune("tsr"),
	}, 3},
}

func TestMinRemoveColumns(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinRemoveColumns(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkMinRemoveColumns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinRemoveColumns(tc.input)
		}
	}
}
