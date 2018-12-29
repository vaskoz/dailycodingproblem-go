package day129

import "testing"

var testcases = []struct {
	n, expected float64
}{
	{9.0, 3.0},
	{36.0, 6.0},
	{81.0, 9.0},
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Sqrt(tc.n); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Sqrt(tc.n)
		}
	}
}
