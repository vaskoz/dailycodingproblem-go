package day319

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	start       Puzzle
	expected    []Puzzle
	expectedErr error
}{
	{
		Puzzle{1, 2, 3, 4, 5, 6, 8, 7, 0},
		nil,
		errNotSolvable,
	},
	{
		Puzzle{1, 0, 3, 5, 2, 6, 4, 7, 8},
		[]Puzzle{
			{1, 0, 3, 5, 2, 6, 4, 7, 8},
			{1, 2, 3, 5, 0, 6, 4, 7, 8},
			{1, 2, 3, 0, 5, 6, 4, 7, 8},
			{1, 2, 3, 4, 5, 6, 0, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 0, 8},
			{1, 2, 3, 4, 5, 6, 7, 8, 0},
		},
		nil,
	},
}

func TestEightPuzzle(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res, err := EightPuzzle(tc.start); !reflect.DeepEqual(res, tc.expected) || err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v). Items %v", tc.expected, tc.expectedErr, res, err, len(res))
		}
	}
}

func BenchmarkEightPuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EightPuzzle(tc.start) // nolint
		}
	}
}
