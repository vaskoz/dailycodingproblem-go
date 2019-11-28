package day374

import "testing"

// nolint
var testcases = []struct {
	input       []int
	expected    int
	expectedErr error
}{
	{[]int{-5, -3, 2, 3}, 2, nil},
	{[]int{2, 3, -3}, 0, errInputNotSorted},
	{[]int{-5, -3, 5, 13}, 0, errNoSuchIndex},
}

func TestLowestFixedPoint(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res, err := LowestFixedPoint(tc.input); res != tc.expected || err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, res, err)
		}
	}
}

func BenchmarkLowestFixedPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LowestFixedPoint(tc.input) // nolint
		}
	}
}
