package day176

import "testing"

var testcases = []struct {
	s1, s2 string
	exists bool
}{
	{"abc", "bcd", true},
	{"foo", "bar", false},
	{"a", "bb", false},
}

func TestDoesMappingExist(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := DoesMappingExist(tc.s1, tc.s2); result != tc.exists {
			t.Errorf("Expected %v got %v", tc.exists, result)
		}
	}
}

func BenchmarkDoesMappingExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DoesMappingExist(tc.s1, tc.s2)
		}
	}
}
