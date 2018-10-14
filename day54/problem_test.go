package sudoku

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	board    Board
	n        int
	expected Board
}{
	{Board{4, 0, 0, 0, 0, 0, 0, 0, 5,
		0, 0, 9, 4, 0, 2, 8, 0, 0,
		0, 6, 0, 0, 5, 0, 0, 9, 0,
		0, 3, 0, 0, 8, 0, 0, 2, 0,
		0, 0, 2, 5, 0, 1, 3, 0, 0,
		0, 9, 0, 0, 4, 0, 0, 7, 0,
		0, 1, 0, 0, 6, 0, 0, 5, 0,
		0, 0, 8, 1, 0, 5, 9, 0, 0,
		5, 0, 0, 0, 0, 0, 0, 0, 7},
		3,
		Board{4, 8, 7, 6, 1, 9, 2, 3, 5,
			3, 5, 9, 4, 7, 2, 8, 1, 6,
			2, 6, 1, 3, 5, 8, 7, 9, 4,
			1, 3, 4, 7, 8, 6, 5, 2, 9,
			6, 7, 2, 5, 9, 1, 3, 4, 8,
			8, 9, 5, 2, 4, 3, 6, 7, 1,
			9, 1, 3, 8, 6, 7, 4, 5, 2,
			7, 4, 8, 1, 2, 5, 9, 6, 3,
			5, 2, 6, 9, 3, 4, 1, 8, 7}},
}

func TestSolver(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, _ := Solver(tc.board, tc.n); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Solver(tc.board, tc.n)
		}
	}
}
