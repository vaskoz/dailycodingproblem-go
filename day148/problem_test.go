package day148

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	bits  int
	codes []string
}{
	{2, []string{"00", "01", "11", "10"}},
	{1, []string{"0", "1"}},
	{0, nil},
	{-10, nil},
	{3, []string{"000", "001", "011", "010", "110", "111", "101", "100"}},
}

func TestGrayCodes(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := GrayCodes(tc.bits); !reflect.DeepEqual(result, tc.codes) {
			t.Errorf("Expected %v got %v", tc.bits, result)
		}
	}
}

func BenchmarkGrayCodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			GrayCodes(tc.bits)
		}
	}
}
