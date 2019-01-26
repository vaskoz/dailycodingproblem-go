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

// MinSquaredDP returns the smallest number
// of squared integers which sum to n.
// Runs in O(N^2) time and O(N) space.
func MinSquaredDP(n int) int {
	table := make([]int, 0, n+1)
	table = append(table, 0, 1, 2, 3)
	for i := 4; i <= n+1; i++ {
		table = append(table, i)
		for x := 1; x < i+1; x++ {
			if v := x * x; v <= i && table[i-v]+1 < table[i] {
				table[i] = table[i-v] + 1
			}
		}
	}
	return table[n]
}
