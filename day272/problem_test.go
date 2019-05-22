package day272

import "testing"

// nolint
var testcases = []struct {
	n, faces, total int
	possibleWays    int
}{
	{3, 6, 7, 15},
}

func TestPossibleWaysToThrow(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := PossibleWaysToThrow(tc.n, tc.faces, tc.total); result != tc.possibleWays {
			t.Errorf("Expected %v, got %v", tc.possibleWays, result)
		}
	}
}

func BenchmarkPossibleWaysToThrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PossibleWaysToThrow(tc.n, tc.faces, tc.total)
		}
	}
}
