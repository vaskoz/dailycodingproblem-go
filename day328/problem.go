package day328

import "math"

// Match represents the IDs of the winner and loser of each match.
type Match struct {
	Winner, Loser int
}

// EloGames returns the new ELO scores for all the players after playing
// the specified matches in order.
func EloGames(scores []float64, kFactor float64, matches []Match) []float64 {
	for _, match := range matches {
		scores[match.Winner], scores[match.Loser] =
			playGame(scores[match.Winner], scores[match.Loser], kFactor)
	}

	return scores
}

func playGame(winner, loser, kFactor float64) (float64, float64) {
	probWinnerWins := probabilityOfWin(loser, winner)
	probLoserWins := probabilityOfWin(winner, loser)
	winner += kFactor * (1 - probWinnerWins)
	loser += kFactor * -probLoserWins

	return winner, loser
}

func probabilityOfWin(loser, winner float64) float64 {
	return 1.0 * 1.0 / (1 + 1.0*math.Pow(10, 1.0*(loser-winner)/400))
}
