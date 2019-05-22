package day273

import "testing"

// nolint
var testcases = []struct {
	sorted      []int
	expected    int
	expectedErr error
}{
	{[]int{-6, 0, 2, 40}, 2, nil},
	{[]int{1, 5, 7, 8}, 0, ErrNoFixedPoint()},
	{[]int{-10, -1, 0, 3, 10, 11, 30, 50, 100}, 3, nil},
}

func TestFindFixedPoint(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := FindFixedPoint(tc.sorted); result != tc.expected || err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func BenchmarkFindFixedPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindFixedPoint(tc.sorted)
		}
	}
}
