package day267

import "testing"

// nolint
var testcases = []struct {
	board   Chessboard
	inCheck bool
}{
	{
		Chessboard{
			"...K....",
			"........",
			".B......",
			"......P.",
			".......R",
			"..N.....",
			"........",
			".....Q..",
		},
		true,
	},
	{
		Chessboard{
			"...K....",
			"........",
			".P..N...",
			"......P.",
			".......R",
			"..N.....",
			"........",
			".....Q..",
		},
		true,
	},
}

func TestIsSoloBlackKingInCheck(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := IsSoloBlackKingInCheck(tc.board); result != tc.inCheck {
			t.Errorf("testcase id:%d expected %v, got %v", tcid, tc.inCheck, result)
		}
	}
}

func BenchmarkIsSoloBlackKingInCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsSoloBlackKingInCheck(tc.board)
		}
	}
}
