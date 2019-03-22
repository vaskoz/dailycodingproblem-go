package day209

import "testing"

var testcases = []struct {
	one, two, three string
	longest         int
}{
	{"epidemiologist", "refrigeration", "supercalifragilisticexpialodocious", 5},
	{"epidemiologist", "supercalifragilisticexpialodocious", "refrigeration", 5},
}

func TestLengthLongestCommonSubseqOf3(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LengthLongestCommonSubseqOf3(tc.one, tc.two, tc.three); tc.longest != result {
			t.Errorf("Expected %v, got %v", tc.longest, result)
		}
	}
}

func BenchmarkLengthLongestCommonSubseqOf3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LengthLongestCommonSubseqOf3(tc.one, tc.two, tc.three)
		}
	}
}
