package day209

import "testing"

var testcases = []struct {
	one, two, three string
	longest         int
}{
	{"epidemiologist", "refrigeration", "supercalifragilisticexpialodocious", 5},
	{"epidemiologist", "supercalifragilisticexpialodocious", "refrigeration", 5},
}

func TestLengthLongestCommonSubseqOf3Recursive(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LengthLongestCommonSubseqOf3Recursive(tc.one, tc.two, tc.three); tc.longest != result {
			t.Errorf("Expected %v, got %v", tc.longest, result)
		}
	}
}

func BenchmarkLengthLongestCommonSubseqOf3Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LengthLongestCommonSubseqOf3Recursive(tc.one, tc.two, tc.three)
		}
	}
}

func TestLengthLongestCommonSubseqOf3Iterative(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LengthLongestCommonSubseqOf3Iterative(tc.one, tc.two, tc.three); tc.longest != result {
			t.Errorf("Expected %v, got %v", tc.longest, result)
		}
	}
}

func BenchmarkLengthLongestCommonSubseqOf3Iterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LengthLongestCommonSubseqOf3Iterative(tc.one, tc.two, tc.three)
		}
	}
}
