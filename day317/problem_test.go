package day317

import "testing"

// nolint
var testcases = []struct {
	m, n, expected uint64
}{
	{1, 10, 0},
	{12, 15, 12},
	{10, 1, 0},
	{15, 12, 12},
	{19, 17, 16},
	{20, 10, 0},
}

func TestBitwiseAndInclusiveRange(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := BitwiseAndInclusiveRange(tc.m, tc.n); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkBitwiseAndInclusiveRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BitwiseAndInclusiveRange(tc.m, tc.n)
		}
	}
}
