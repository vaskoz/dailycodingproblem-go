package day238

import (
	"testing"
)

var testcases = []struct {
	deck     []Card
	mostWins int
}{
	{
		[]Card{
			{"clubs", "ace"}, {"spades", "six"}, {"hearts", "queen"}, {"clubs", "eight"}, {"spades", "five"}, {"hearts", "jack"}, {"hearts", "ace"}, {"clubs", "three"}, {"diamonds", "five"}, {"clubs", "two"}, {"diamonds", "queen"}, {"diamonds", "four"}, {"hearts", "five"}, {"hearts", "six"}, {"diamonds", "jack"}, {"hearts", "seven"}, {"clubs", "five"}, {"spades", "seven"}, {"clubs", "ten"}, {"diamonds", "three"}, {"spades", "eight"}, {"diamonds", "ten"}, {"diamonds", "two"}, {"spades", "nine"}, {"spades", "two"}, {"spades", "jack"}, {"diamonds", "eight"}, {"spades", "three"}, {"spades", "four"}, {"clubs", "four"}, {"clubs", "seven"}, {"diamonds", "nine"}, {"diamonds", "seven"}, {"diamonds", "king"}, {"hearts", "three"}, {"hearts", "eight"}, {"hearts", "two"}, {"clubs", "jack"}, {"spades", "ace"}, {"spades", "king"}, {"hearts", "ten"}, {"clubs", "king"}, {"clubs", "six"}, {"diamonds", "ace"}, {"hearts", "nine"}, {"spades", "ten"}, {"spades", "queen"}, {"hearts", "four"}, {"clubs", "nine"}, {"diamonds", "six"}, {"clubs", "queen"}, {"hearts", "king"},
		},
		8,
	},
	{
		[]Card{
			{"clubs", "ace"}, {"clubs", "two"}, {"clubs", "king"}, {"spades", "ten"}, {"hearts", "three"}, {"hearts", "six"}, {"clubs", "eight"}, {"spades", "nine"}, {"hearts", "ace"}, {"clubs", "six"}, {"clubs", "nine"}, {"hearts", "eight"}, {"spades", "three"}, {"clubs", "three"}, {"diamonds", "jack"}, {"hearts", "queen"}, {"diamonds", "nine"}, {"diamonds", "three"}, {"hearts", "four"}, {"hearts", "seven"}, {"spades", "four"}, {"hearts", "two"}, {"spades", "jack"}, {"spades", "ace"}, {"diamonds", "four"}, {"spades", "two"}, {"diamonds", "queen"}, {"hearts", "five"}, {"diamonds", "ace"}, {"spades", "five"}, {"clubs", "queen"}, {"hearts", "ten"}, {"diamonds", "king"}, {"clubs", "five"}, {"hearts", "king"}, {"diamonds", "eight"}, {"spades", "queen"}, {"diamonds", "five"}, {"clubs", "ten"}, {"spades", "six"}, {"diamonds", "two"}, {"spades", "eight"}, {"diamonds", "seven"}, {"clubs", "jack"}, {"hearts", "jack"}, {"spades", "seven"}, {"diamonds", "six"}, {"clubs", "four"}, {"spades", "king"}, {"diamonds", "ten"}, {"hearts", "nine"}, {"clubs", "seven"},
		},
		5,
	},
	{
		[]Card{
			{"clubs", "four"}, {"clubs", "six"}, {"diamonds", "five"}, {"hearts", "three"}, {"clubs", "eight"}, {"spades", "ten"}, {"diamonds", "eight"}, {"hearts", "six"}, {"clubs", "two"}, {"hearts", "queen"}, {"spades", "queen"}, {"spades", "six"}, {"spades", "ace"}, {"diamonds", "ten"}, {"spades", "nine"}, {"hearts", "king"}, {"clubs", "queen"}, {"hearts", "four"}, {"diamonds", "six"}, {"spades", "four"}, {"clubs", "ace"}, {"clubs", "jack"}, {"diamonds", "nine"}, {"diamonds", "king"}, {"hearts", "five"}, {"hearts", "jack"}, {"diamonds", "three"}, {"spades", "five"}, {"clubs", "five"}, {"spades", "seven"}, {"spades", "two"}, {"hearts", "ten"}, {"spades", "king"}, {"hearts", "eight"}, {"diamonds", "two"}, {"hearts", "ace"}, {"clubs", "three"}, {"hearts", "nine"}, {"hearts", "seven"}, {"hearts", "two"}, {"spades", "jack"}, {"spades", "three"}, {"diamonds", "jack"}, {"diamonds", "ace"}, {"spades", "eight"}, {"clubs", "seven"}, {"diamonds", "four"}, {"diamonds", "queen"}, {"clubs", "ten"}, {"diamonds", "seven"}, {"clubs", "king"}, {"clubs", "nine"},
		},
		6,
	},
	{
		[]Card{
			{"diamonds", "five"}, {"spades", "ten"}, {"spades", "king"}, {"clubs", "queen"}, {"hearts", "ten"}, {"spades", "jack"}, {"hearts", "four"}, {"spades", "two"}, {"clubs", "seven"}, {"hearts", "eight"}, {"spades", "four"}, {"hearts", "king"}, {"hearts", "six"}, {"hearts", "nine"}, {"spades", "six"}, {"clubs", "four"}, {"diamonds", "nine"}, {"spades", "queen"}, {"diamonds", "four"}, {"hearts", "five"}, {"clubs", "two"}, {"spades", "ace"}, {"spades", "three"}, {"spades", "nine"}, {"clubs", "six"}, {"diamonds", "king"}, {"clubs", "ten"}, {"clubs", "jack"}, {"hearts", "two"}, {"hearts", "three"}, {"spades", "five"}, {"diamonds", "ten"}, {"clubs", "nine"}, {"clubs", "three"}, {"diamonds", "queen"}, {"diamonds", "seven"}, {"hearts", "jack"}, {"hearts", "ace"}, {"hearts", "seven"}, {"clubs", "ace"}, {"diamonds", "jack"}, {"spades", "eight"}, {"diamonds", "ace"}, {"hearts", "queen"}, {"diamonds", "three"}, {"spades", "seven"}, {"diamonds", "two"}, {"clubs", "king"}, {"diamonds", "eight"}, {"clubs", "five"}, {"clubs", "eight"}, {"diamonds", "six"},
		},
		8,
	},
}

func TestMostPlayerWins(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if most := MostPlayerWins(tc.deck); most != tc.mostWins {
			t.Errorf("Expected %v, got %v", tc.mostWins, most)
		}
	}
}

func BenchmarkMostPlayerWins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MostPlayerWins(tc.deck)
		}
	}
}
