package day66

import (
	"math/rand"
	"time"
)

// BiasedCoin represents a biased coin. Not 50-50.
type BiasedCoin struct {
	zeroBias int
	r        *rand.Rand
}

// Toss throws the coin and returns 0 or 1
func (bc *BiasedCoin) Toss() int {
	if val := bc.r.Intn(100); val < bc.zeroBias {
		return 0
	}
	return 1
}

// NewBiasedCoin returns a biased coin for flipping.
// zeroBias represents the amount of bias towards a zero value.
// The higher the value of zeroBias the more you will see zeros.
// a zeroBias of 50 will be a fair coin.
func NewBiasedCoin(zeroBias int) *BiasedCoin {
	return &BiasedCoin{zeroBias: zeroBias, r: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// FairTosser represents a fair 50-50 coin tosser even if given an unfair coin.
type FairTosser struct {
	bc *BiasedCoin
}

// Toss throws the coin and returns 0 or 1 and should be a 50-50 chance.
func (ft *FairTosser) Toss() int {
	firstToss := ft.bc.Toss()
	secondToss := ft.bc.Toss()
	for firstToss == secondToss {
		firstToss = ft.bc.Toss()
		secondToss = ft.bc.Toss()
	}
	return firstToss
}

// NewFairTosser makes a biased coin fair.
func NewFairTosser(bc *BiasedCoin) *FairTosser {
	return &FairTosser{bc: bc}
}
