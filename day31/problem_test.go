package day31

import "testing"

var testcases = []struct {
	s1, s2 string
	diff   int
}{
	{"kitten", "sitting", 3},
}

func TestEditDistance(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if diff := EditDistance(tc.s1, tc.s2); diff != tc.diff {
			t.Errorf("Expected %v got %v", tc.diff, diff)
		}
	}
}

func BenchmarkEditDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EditDistance(tc.s1, tc.s2)
		}
	}
}
