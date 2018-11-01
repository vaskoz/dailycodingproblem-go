package day71

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

// Rand7 returns a random integer between 1 and 7 inclusive.
func Rand7() int {
	return GetRandom().Intn(7) + 1
}

// Rand5 returns a random integer between 1 and 5 inclusive.
// Using only Rand7 as a randomization source.
func Rand5() int {
	result := 7
	for result > 5 {
		result = Rand7()
	}
	return result
}
