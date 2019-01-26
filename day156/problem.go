package day156

import (
	"math"
)

// MinSquared returns the smallest number
// of squared integers which sum to n.
// Runs in exponential time.
func MinSquared(n int) int {
	return minSquared(n, 0, n)
}

func minSquared(n, level, minSoFar int) int {
	if n <= 3 {
		return n
	}
	end := int(math.Sqrt(float64(n))) + 1
	for i := end; i > 0; i-- {
		if x := i * i; x <= n {
			if y := 1 + minSquared(n-x, level+1, minSoFar); y < minSoFar {
				minSoFar = y
			}
		}
	}
	return minSoFar
}
