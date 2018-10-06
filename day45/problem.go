package day45

import (
	"math/rand"
	"time"
)

var random *rand.Rand

// GetRandom returns a *math.Rand.
func GetRandom() *rand.Rand {
	if random == nil {
		random = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return random
}

// Rand5 returns a random integer between 1 and 5 inclusive with uniform probability.
func Rand5() int {
	return GetRandom().Intn(5) + 1
}

// Rand7 returns a random integer between 1 and 7 inclusive with uniform probability.
// The trick is using Rand5 to accomplish that.
func Rand7() int {
	twentyOne := 22
	for twentyOne > 21 {
		twentyOne = 5*Rand5() + Rand5() - 5
	}
	return (twentyOne % 7) + 1
}
