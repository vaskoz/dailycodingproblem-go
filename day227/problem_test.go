package day227

import "testing"

var testcases = []struct {
	board Board
	dict  []string
	valid []string
}{
	{Board{
		[]rune("aaga"),
		[]rune("abaf"),
		[]rune("acea"),
		[]rune("adaa"),
	},
		[]string{"abcdefga", "abcda", "foobar"},
		[]string{"abcdefga", "abcda"},
	},
}

func TestBoggleSolver(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		words := BoggleSolver(tc.board, tc.dict)
		m := make(map[string]struct{}, len(tc.valid))
		for _, w := range tc.valid {
			m[w] = struct{}{}
		}
		for _, w := range words {
			if _, found := m[w]; !found {
				t.Errorf("got word %v, but didn't expect it", w)
			} else {
				delete(m, w)
			}
		}
		if len(m) != 0 {
			t.Errorf("expected more words: %v", m)
		}
	}
}
