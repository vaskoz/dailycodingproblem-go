package day460

import "testing"

// nolint
var testcases = []struct {
	input string
	flips int
}{
	{"xyxxxyxyy", 2},
	{"yxxyyyy", 1},
	{"yyyxy", 1},
	{"yxyxyx", 3},
	{"yyyyxyyyxyyyyxx", 4},
	{"xxxxxxx", 0},
	{"yyyyyyyyyy", 0},
	{"xyyyyyyyy", 0},
	{"xxxxxxxxyyy", 0},
	{"xxxyyxxyyxxyyyyyyyyy", 4},
}

func TestFlipsXsBeforeYs(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if flips := FlipsXsBeforeYs(tc.input); flips != tc.flips {
			t.Errorf("Expected %v, got %v", tc.flips, flips)
		}
	}
}

func BenchmarkFlipsXsBeforeYs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FlipsXsBeforeYs(tc.input)
		}
	}
}
