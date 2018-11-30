package day98

import "testing"

var testcases = []struct {
	board    Board
	word     string
	expected bool
}{
	{Board{
		[]rune("ABCE"),
		[]rune("SFCS"),
		[]rune("ADEE"),
	}, "ABCCED", true},
	{Board{
		[]rune("ABCE"),
		[]rune("SFCS"),
		[]rune("ADEE"),
	}, "SEE", true},
	{Board{
		[]rune("ABCE"),
		[]rune("SFCS"),
		[]rune("ADEE"),
	}, "ABCB", false},
	{Board{
		[]rune("ABCE"),
		[]rune("SFCS"),
		[]rune("ADEE"),
	}, "ABCESCFSADEE", true},
	{Board{
		[]rune("ABCE"),
		[]rune("SFCS"),
		[]rune("ADEE"),
	}, "ABCESCFSADEEA", false},
	{Board{
		[]rune("ABCE"),
		[]rune("SFCS"),
		[]rune("ADEE"),
	}, "BCEA", true},
}

func TestDoesExistSeqAdj(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := DoesExistSeqAdj(tc.board, tc.word); result != tc.expected {
			t.Errorf("TC%d expected %v got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkDoesExistSeqAdj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DoesExistSeqAdj(tc.board, tc.word)
		}
	}
}
