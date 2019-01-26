package day156

import "testing"

var testcases = []struct {
	n, expected int
}{
	{13, 2},
	{27, 3},
}

func TestMinSquared(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinSquared(tc.n); result != tc.expected {
			t.Errorf("For N=%d Expected %v got %v", tc.n, tc.expected, result)
		}
	}
}

func BenchmarkMinSquared(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinSquared(tc.n)
		}
	}
}

func TestMinSquaredDP(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinSquaredDP(tc.n); result != tc.expected {
			t.Errorf("For N=%d Expected %v got %v", tc.n, tc.expected, result)
		}
	}
}

func BenchmarkMinSquaredDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinSquaredDP(tc.n)
		}
	}
}
