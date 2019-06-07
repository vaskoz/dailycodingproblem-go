package day288

import "testing"

// nolint
var testcases = []struct {
	num        int
	iterations int
	err        error
}{
	{1234, 3, nil},
	{0, 0, ErrInvalidInput()},
	{1111, 0, ErrInvalidInput()},
	{4352, 3, nil},
	{9999, 0, ErrInvalidInput()},
	{12, 3, nil},
	{2111, 5, nil},
	{104, 7, nil},
}

func TestKaprekarConstantIterations(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if iterations, err := KaprekarConstantIterations(tc.num); iterations != tc.iterations || err != tc.err {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.iterations, tc.err, iterations, err)
		}
	}
}

func BenchmarkKaprekarConstantIterations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			KaprekarConstantIterations(tc.num) // nolint
		}
	}
}
