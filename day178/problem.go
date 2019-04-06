package day178

import (
	"math/rand"
	"time"
)

// Dice represents a 6-sided random die.
type Dice struct {
	r *rand.Rand
}

// Next returns the value of the next die throw.
func (d *Dice) Next() int {
	if d.r == nil {
		d.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return d.r.Intn(6) + 1
}

// FirstGame returns the number of rolls required to meet the condition.
// Condition is a five followed by a six.
func FirstGame(d *Dice) int {
	var rolls int
	var prevFive bool
	for {
		rolls++
		value := d.Next()
		switch {
		case value == 5:
			prevFive = true
		case value == 6 && prevFive:
			return rolls
		default:
			prevFive = false
		}
	}
}

// SecondGame returns the number of rolls required to meet the condition.
// Condition is a five followed by another five.
func SecondGame(d *Dice) int {
	var rolls int
	var prevFive bool
	for {
		rolls++
		value := d.Next()
		switch {
		case value == 5 && prevFive:
			return rolls
		case value == 5:
			prevFive = true
		default:
			prevFive = false
		}
	}
}
