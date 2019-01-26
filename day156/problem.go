package day156

import (
	"math"
)

// MinSquared returns the smallest number
// of squared integers which sum to n.
// Runs in exponential time.
func MinSquared(n int) int {
	if n <= 3 {
		return n
	}
	result := n
	end := int(math.Sqrt(float64(n))) + 1
	for i := end; i > 0; i-- {
		if x := i * i; x <= n {
			if y := 1 + MinSquared(n-x); y < result {
				result = y
			}
		}
	}
	return result
}
