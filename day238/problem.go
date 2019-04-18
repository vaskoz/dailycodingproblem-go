package day238

// Suit represents the 4 suits: Hearts, Spades, Diamonds and Clubs.
type Suit string

// Rank represents the face of the card.
type Rank string

// Card represents a playing card with a Suit and Rank.
type Card struct {
	s Suit
	r Rank
}

var cardValues = map[Rank]int{
	"ace":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

func (c *Card) value() int {
	return cardValues[c.r]
}

// MostPlayerWins returns the most number of possible player wins.
// Dealer gets the first 2 cards of each hand.
// Player gets second 2 cards.
// Player tries every game based on drawing cards up until bust.
// Dealer draws until greater than or equal to 17 and stops.
func MostPlayerWins(deck []Card) int {
	if len(deck) < 4 { // not enough cards to deal
		return 0
	}
	dealer := deck[0].value() + deck[1].value()
	player := deck[2].value()
	var most int
	start := 3
	cards := 1
	for {
		playerHand := player
		dealerHand := dealer
		next := start
		for i := 0; i < cards; i++ {
			if next == len(deck) {
				return 0
			}
			playerHand += deck[next].value()
			next++
		}
		if playerHand > 21 {
			break
		}
		for dealerHand < 17 {
			if next == len(deck) {
				return 0
			}
			dealerHand += deck[next].value()
			next++
		}
		if dealerHand > 21 || playerHand > dealerHand {
			most = max(most, 1+MostPlayerWins(deck[next:]))
		} else {
			most = max(most, MostPlayerWins(deck[next:]))
		}
		cards++
	}
	return most
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
