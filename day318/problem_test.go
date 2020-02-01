package day318

import "testing"

// nolint
var testcases = []struct {
	N, M, B  int
	expected int
}{
	{4, 2, 1, 2},
	{4, 2, 0, 16},
	{4, 2, 2, 0},
	{9, 2, 1, 2},
	{9, 2, 0, 512},
	{9, 3, 2, 6},
	{9, 3, 1, 768},
	{9, 3, 0, 19683},
}

func TestNumberOfValidPlayListsRec(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := NumberOfValidPlayListsRec(tc.N, tc.M, tc.B); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkNumberOfValidPlayListsRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumberOfValidPlayListsRec(tc.N, tc.M, tc.B)
		}
	}
}

func TestNumberOfValidPlayListsIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := NumberOfValidPlayListsIterative(tc.N, tc.M, tc.B); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkNumberOfValidPlayListsIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumberOfValidPlayListsIterative(tc.N, tc.M, tc.B)
		}
	}
}
