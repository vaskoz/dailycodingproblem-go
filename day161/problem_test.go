package day161

import "testing"

var testcases = []struct {
	given, expected uint32
}{
	{4042322160, 252645135},
	{3, 3221225472},
}

func TestReverse32Bits(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Reverse32Bits(tc.given); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkReverse32Bits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Reverse32Bits(tc.given)
		}
	}
}
