package day88

import "testing"

var testcases = []struct {
	dividend, divisor, quotient int32
	expectedError               error
}{
	{100, 10, 10, nil},
	{100, -10, -10, nil},
	{100, 0, 0, ErrDivideByZero()},
	{1000000, 1, 1000000, nil},
}

func TestBruteForceDivision(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := BruteForceDivision(tc.dividend, tc.divisor); result != tc.quotient || err != tc.expectedError {
			t.Errorf("Expected %v,%v got %v,%v", tc.quotient, tc.expectedError, result, err)
		}
	}
}

func BenchmarkBruteForceDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BruteForceDivision(tc.dividend, tc.divisor)
		}
	}
}

func TestBitShiftDivision(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := BitShiftDivision(tc.dividend, tc.divisor); result != tc.quotient || err != tc.expectedError {
			t.Errorf("Expected %v,%v got %v,%v", tc.quotient, tc.expectedError, result, err)
		}
	}
}

func BenchmarkBitShiftDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BitShiftDivision(tc.dividend, tc.divisor)
		}
	}
}
