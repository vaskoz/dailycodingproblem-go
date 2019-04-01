package day222

import "testing"

var testcases = []struct {
	path, expected string
}{
	{"/usr/bin/../bin/./scripts/../", "/usr/bin/"},
	{"/usr/bin/../bin/./scripts/../foobar.txt", "/usr/bin/foobar.txt"},
	{"/usr/bin/../../bin/./scripts/foobar.txt", "/bin/scripts/foobar.txt"},
}

func TestShortestStandardizedPath(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ShortestStandardizedPath(tc.path); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkShortestStandardizedPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ShortestStandardizedPath(tc.path)
		}
	}
}
