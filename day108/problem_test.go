package day108

import "testing"

var testcases = []struct {
	A, B     string
	expected bool
}{
	{"abcde", "cdeab", true},
	{"abc", "acb", false},
	{"abc", "ac", false},
	{"", "", true},
	{"abcab", "cabab", true},
	{"abcab", "acbab", false},
}

func TestEqualWithShifting(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := EqualWithShifting(tc.A, tc.B); result != tc.expected {
			t.Errorf("For TCID:%d expected %v got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkEqualWithShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EqualWithShifting(tc.A, tc.B)
		}
	}
}
