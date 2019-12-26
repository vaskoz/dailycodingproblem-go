package day336

import "testing"

// nolint
var testcases = []struct {
	n, expected int
}{
	{10, 3360},
	{3, 2},
}

func TestPossibleMaxHeaps(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := PossibleMaxHeaps(tc.n); res != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkPossibleMaxHeaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PossibleMaxHeaps(tc.n)
		}
	}
}
