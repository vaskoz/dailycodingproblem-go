package day138

import "testing"

var testcases = []struct {
	denom    []int
	n        int
	expected int
	err      error
}{
	{[]int{1, 5, 10, 25}, 16, 3, nil},
	{[]int{25}, 16, 0, ErrNotPossible()},
	{[]int{2, 5}, 16, 5, nil},
	{[]int{2, 3, 6}, 6, 1, nil},
}

func TestMinCurrencyRequiredBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := MinCurrencyRequiredBrute(tc.n, tc.denom); err != tc.err || result != tc.expected {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.expected, tc.err, result, err)
		}
	}
}

func BenchmarkMinCurrencyRequiredBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinCurrencyRequiredBrute(tc.n, tc.denom)
		}
	}
}

func TestMinCurrencyRequiredSmarter(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := MinCurrencyRequiredSmarter(tc.n, tc.denom); err != tc.err || result != tc.expected {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.expected, tc.err, result, err)
		}
	}
}

func BenchmarkMinCurrencyRequiredSmarter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinCurrencyRequiredSmarter(tc.n, tc.denom)
		}
	}
}
