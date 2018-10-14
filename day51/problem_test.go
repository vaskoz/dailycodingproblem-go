package day51

import "testing"

func TestShuffle(t *testing.T) {
	t.Parallel()
	cards := NewDeck()
	Shuffle(cards, 5)
}

func TestNewDeck(t *testing.T) {
	t.Parallel()
	suits := []Suit{"clubs", "diamonds", "spades", "hearts"}
	ranks := []Rank{"two", "three", "four", "five", "six", "seven",
		"eight", "nine", "ten", "jack", "queen", "king", "ace"}
	index := 0
	cards := NewDeck()
	for _, suit := range suits {
		for _, rank := range ranks {
			if expected := (Card{suit, rank}); cards[index] != expected {
				t.Errorf("Expected %v, got %v", expected, cards[index])
			}
			index++
		}
	}
}

func BenchmarkNewDeck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDeck()
	}
}
