package day152

import (
	"math/rand"
	"time"
)

// Random returns a random value based on provided probabilities.
type Random interface {
	Get() interface{}
}

type randomDistribution struct {
	r    *rand.Rand
	prob []float64
	vals []interface{}
}

// NewRandom returns a new instance of Random.
func NewRandom(vals []interface{}, prob []float64) Random {
	newprob := make([]float64, len(prob))
	newprob[0] = prob[0]
	for i := 1; i < len(prob); i++ {
		newprob[i] = newprob[i-1] + prob[i]
	}
	return &randomDistribution{
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
		prob: newprob,
		vals: vals,
	}
}

func (rd *randomDistribution) Get() interface{} {
	r := rd.r.Float64()
	prev := 0.0
	ansI := 0
	for i, p := range rd.prob {
		if r >= prev && r <= p {
			ansI = i
			break
		}
		prev = p
	}
	return rd.vals[ansI]
}
