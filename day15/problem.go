package day15

import (
	"math/rand"
	"time"
)

var randomizer *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// UniformlyRandom reads from a stream of anything and
// returns an element with uniform probability.
// O(N) time and O(1) space
func UniformlyRandom(stream <-chan interface{}) interface{} {
	var result interface{}
	var count int
	for element := range stream {
		if count == 0 {
			result = element
		} else if randomizer.Intn(count) == 0 {
			result = element
		}
		count++
	}
	return result
}
