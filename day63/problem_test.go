package day63

import "testing"

var testcases = []struct {
	puzzle   []string
	target   string
	expected bool
}{
	{[]string{"FACI", "OBQP", "ANOB", "MASS"}, "FOAM", true},
	{[]string{"FACI", "OBQP", "ANOB", "MASS"}, "MASS", true},
	{[]string{"FACI", "OBQP", "ANOB", "MASS"}, "FBOS", false},
	{[]string{"FACI", "OBQP", "ANOB", "MASS"}, "F", true},
	{[]string{"FACI", "OBQP", "ANOB", "MASS"}, "MASSS", false},
}

func TestFindWordLD(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		puzzle := make([][]rune, len(tc.puzzle))
		for i := range puzzle {
			puzzle[i] = []rune(tc.puzzle[i])
		}
		if result := FindWordLD(puzzle, []rune(tc.target)); result != tc.expected {
			t.Errorf("TC%d Expected %v got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkFindWordLD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			puzzle := make([][]rune, len(tc.puzzle))
			for x := range puzzle {
				puzzle[x] = []rune(tc.puzzle[x])
			}
			FindWordLD(puzzle, []rune(tc.target))
		}
	}
}
