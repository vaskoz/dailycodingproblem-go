package day455

import "testing"

// nolint
var testcases = []struct {
	board    string
	steps    int
	expected string
}{
	{"*\n*\n*", 1, "***\n"},
	{"*\n*\n*", 2, "*\n*\n*\n"},
	{"**\n**", 2, "**\n**\n"},
	{"**\n*.", 1, "**\n**\n"},
	{"..*\n.*.\n*..", 1, "*\n"},
	{"**.\n*.*\n.*.", 2, "**.\n*.*\n.*.\n"},
	{"**.\n*.*\n.*.", 20, "**.\n*.*\n.*.\n"},
}

func TestGameOfLife(t *testing.T) {
	t.Parallel()

	for n, tc := range testcases {
		gol := NewGameOfLife(tc.board)

		for i := 0; i < tc.steps; i++ {
			gol.Step()
		}

		if result := gol.String(); result != tc.expected {
			t.Errorf("TC%d Expected %v got %v", n, tc.expected, result)
		}
	}
}

func BenchmarkGameOfLife(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			gol := NewGameOfLife(tc.board)
			gol.Step()
			_ = gol.String()
		}
	}
}
