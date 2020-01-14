package day419

import "math"

// SmallestStepsToOne returns the smallest possible
// number of steps to reach run using two rules.
// First, n-1
// Second, a*b=N goto the larger of a and b.
// This is effectively BFS.
func SmallestStepsToOne(n int) int {
	if n < 1 {
		panic("n must be greater than 0")
	}

	seen := make(map[int]struct{})
	q := []int{n}

	var steps int

	for len(q) != 0 {
		nextQ := make([]int, 0, 2*len(q))

		for _, v := range q {
			if v == 1 {
				nextQ = nil
				break
			} else if _, found := seen[v]; !found {
				seen[v] = struct{}{}
				nextQ = append(nextQ, v-1)
				a := v
				for b := int(math.Sqrt(float64(v))); b > 0; b-- {
					if v%b == 0 {
						a = v / b
						break
					}
				}
				nextQ = append(nextQ, a)
			}
		}

		q = nextQ
		steps++
	}

	return steps - 1
}
