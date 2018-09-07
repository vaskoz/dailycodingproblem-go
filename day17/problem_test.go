package day17

import "testing"

var testcases = []struct {
	fs       string
	expected int
}{
	{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
	{"root\n\tsubdir1\n\t\tfile1.ext", 22},
}

func TestFindLongestAbsolutePathLength(t *testing.T) {
	for _, tc := range testcases {
		if result := FindLongestAbsolutePathLength(tc.fs); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkFindLongestAbsolutePathLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindLongestAbsolutePathLength(tc.fs)
		}
	}
}
