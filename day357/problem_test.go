package day357

import "testing"

// nolint
var testcases = []struct {
	tree  string
	depth int
}{
	{"(00)", 1},
	{"((00)(00))", 2},
	{"((((00)0)0)0)", 4},
	{"((0((00)0))(00))", 4},
	{"(0(0(0(0(0(00))))))", 6},
	{"(((0(((00)(00))((00)(00))))0)(0((((0(00))0)0)(00))))", 7},
}

func TestStringTreeDepth(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if depth := StringTreeDepth(tc.tree); depth != tc.depth {
			t.Errorf("For tree %s, expected %v, got %v", tc.tree, tc.depth, depth)
		}
	}
}

func BenchmarkStringTreeDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			StringTreeDepth(tc.tree)
		}
	}
}
