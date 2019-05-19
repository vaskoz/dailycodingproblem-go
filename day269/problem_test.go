package day269

import "testing"

// nolint
var testcases = []struct {
	start, end string
}{
	{".L.R....L", "LL.RRRLLL"},
	{"..R...L.L", "..RR.LLLL"},
}

func TestSimulateDominoes(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SimulateDominoes(tc.start); result != tc.end {
			t.Errorf("Given %s, expected %s, got %s", tc.start, tc.end, result)
		}
	}
}

func BenchmarkSimulateDominoes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SimulateDominoes(tc.start)
		}
	}
}
