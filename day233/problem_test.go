package day233

import "testing"

var testcases = []struct {
	n, fib uint
	err    error
}{
	{0, 0, nil},
	{1, 1, nil},
	{2, 1, nil},
	{12, 144, nil},
	{28, 317811, nil},
	{44, 701408733, nil},
	{87, 679891637638612258, nil},
	{100, 0, ErrOverflow()},
}

func TestFastFibonnaci(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := FastFibonnaci(tc.n); result != tc.fib || err != tc.err {
			t.Errorf("Given n=%d, expected (%d,%v), got (%d,%v)", tc.n, tc.fib, tc.err, result, err)
		}
	}
}

func BenchmarkFastFibonnaci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FastFibonnaci(tc.n) //nolint
		}
	}
}
