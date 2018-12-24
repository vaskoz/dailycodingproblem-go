package day124

import "math"

// FairCoinFlipDownToOne answers how many rounds
// are necessary to reach one coin.
// You put aside tails, and flip all heads on the next round.
// How many rounds to reach just 1 coin.
func FairCoinFlipDownToOne(coins int) int {
	return int(math.Ceil(math.Log2(float64(coins))))
}
