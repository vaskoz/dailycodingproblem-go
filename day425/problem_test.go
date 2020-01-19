package day425

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
	{
		Chessboard{
			"...K....",
			"..P.....",
			".P..N...",
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
			".P......",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"...R.Q..",
		},
		true,
	},
	{
		Chessboard{
			"K.......",
			"........",
			".P......",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"...R...Q",
		},
		true,
	},
	{
		Chessboard{
			"Q.......",
			"........",
			".PK.....",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"...R....",
		},
		true,
	},
	{
		Chessboard{
			"....Q...",
			"........",
			".PK.....",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"...R....",
		},
		true,
	},
	{
		Chessboard{
			"........",
			"........",
			".PK...R.",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"........",
		},
		true,
	},
	{
		Chessboard{
			"........",
			"........",
			"Q.K.....",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"........",
		},
		true,
	},
	{
		Chessboard{
			"..R.....",
			"........",
			"..K.....",
			"......P.",
			".......R",
			"..N.....",
			"........",
			"........",
		},
		true,
	},
	{
		Chessboard{
			"...K....",
			"........",
			"........",
			"PPPPPPPP",
			".B.....R",
			"P.N.....",
			"....Q...",
			".R......",
		},
		false,
	},
	{ // missing king
		Chessboard{
			"........",
			"........",
			"........",
			"PPPPPPPP",
			".B.....R",
			"P.N.....",
			"....Q...",
			".R......",
		},
		false,
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
