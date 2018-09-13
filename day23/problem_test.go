package day23

import "testing"

var testcases = []struct {
	board      Board
	start, end Cell
	expected   int
	err        error
}{
	{Board{
		{false, false, false, false},
		{true, true, false, true},
		{false, false, false, true},
		{false, false, false, false},
	}, Cell{3, 0}, Cell{0, 0}, 7, nil},
	{Board{
		{false, true, false, false},
		{true, true, false, true},
		{false, false, false, true},
		{false, false, false, false},
	}, Cell{3, 0}, Cell{0, 0}, 0, errNoPath},
}

func TestMinimumSteps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := MinimumSteps(tc.board, tc.start, tc.end,
			make(map[Cell]struct{})); result != tc.expected || err != tc.err {
			t.Errorf("Expected (%v, %v) but got (%v, %v)",
				tc.expected, tc.err, result, err)
		}
	}
}

func BenchmarkMinimumSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumSteps(tc.board, tc.start, tc.end, make(map[Cell]struct{}))
		}
	}
}
