package day419

import "testing"

// nolint
var testcases = []struct {
	n, expected int
}{
	{100, 5},
	{1, 0}, // 1
	{2, 1}, // 2 -> 1
	{3, 2}, // 3 -> 2 -> 1
	{4, 2}, // 4 -> 2 -> 1
	{5, 3}, // 5 -> 4 -> 2 -> 1
	{6, 3}, // 6 -> 3 -> 2 -> 1
}

func TestSmallestStepsToOne(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := SmallestStepsToOne(tc.n); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestSmallestStepsToOneBadInput(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	}()
	SmallestStepsToOne(0)
}

func BenchmarkSmallestStepsToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestStepsToOne(tc.n)
		}
	}
}
