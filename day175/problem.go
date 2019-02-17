package day175

import (
	"math/rand"
	"time"
)

// State represents states in the Markov chain.
type State string

// Transition represents the probabilities of transitions.
type Transition map[State]map[State]float64

type probabilities map[State][]threshold

type threshold struct {
	s State
	p float64
}

// RunMarkovChain runs the Markov chain from the start state for numSteps.
// The result is the number of times each state occurred.
func RunMarkovChain(start State, states Transition, numSteps int) map[State]int {
	result := make(map[State]int)
	probs := buildProbabilities(states)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numSteps; i++ {
		p := r.Float64()
		d := probs[start]
		var last float64
		for _, t := range d {
			if p > last && p < t.p {
				result[t.s]++
				start = t.s
				break
			}
			last = t.p
		}
	}
	return result
}

func buildProbabilities(states Transition) probabilities {
	result := make(probabilities)
	for s1, sub := range states {
		var data []threshold
		var p float64
		for s2, fl := range sub {
			p += fl
			data = append(data, threshold{s2, p})
		}
		result[s1] = data
	}
	return result
}
