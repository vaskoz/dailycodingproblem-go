package day420

import "testing"

// nolint
var testcases = []struct {
	n, expected int
}{
	{1, 19},
	{2, 28},
	{10, 109},
}

func TestPerfectNth(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := PerfectNth(tc.n); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func TestPerfectNthPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for an argument less than 1")
		}
	}()
	PerfectNth(0)
}

func BenchmarkPerfectNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PerfectNth(tc.n)
		}
	}
}
