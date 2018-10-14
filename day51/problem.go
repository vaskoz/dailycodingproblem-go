package day51

import (
	"math/rand"
	"time"
)

// Suit represents the 4 suits: Hearts, Spades, Diamonds and Clubs.
type Suit string

// Rank represents the face of the card.
type Rank string

// Card represents a playing card with a Suit and Rank.
type Card struct {
	s Suit
	r Rank
}

// Shuffle randomly shuffles the cards using a random function
// that returns values between 1 and k inclusive.
func Shuffle(cards []Card, k int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

// NewDeck returns a deterministic deck of cards.
func NewDeck() []Card {
	suits := []Suit{"clubs", "diamonds", "spades", "hearts"}
	ranks := []Rank{"two", "three", "four", "five", "six", "seven",
		"eight", "nine", "ten", "jack", "queen", "king", "ace"}
	cards := make([]Card, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{suit, rank})
		}
	}
	return cards
}
