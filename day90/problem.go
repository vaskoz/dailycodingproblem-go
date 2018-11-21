package day90

import (
	"math/rand"
	"time"
)

var random *rand.Rand

// getRandom returns a *math.Rand.
func getRandom() *rand.Rand {
	if random == nil {
		random = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return random
}

// RandMissingNumbers randomly generates a number from 0 to n-1 that isn't in l.
// Runs in O(N) time and O(l) space.
func RandMissingNumbers(n int, l []int) int {
	m := make(map[int]struct{})
	for _, v := range l {
		m[v] = struct{}{}
	}
	size := n - len(m)
	if size == 0 {
		return -1
	}
	pos := getRandom().Intn(size)
	val := -1
	var i int
	for i = 0; i < n; i++ {
		if _, found := m[i]; !found {
			val++
		}
		if val == pos {
			break
		}
	}
	return i
}
