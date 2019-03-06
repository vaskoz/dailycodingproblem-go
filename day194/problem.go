package day194

// CountIntersectionsBrute does a brute force check for intersections.
// Runs in O(N^2) time.
func CountIntersectionsBrute(p, q []int) int {
	intersections := 0
	for i := range p {
		for j := i + 1; j < len(q); j++ {
			diffP := p[j] - p[i]
			diffQ := q[j] - q[i]
			if (diffP < 0 && diffQ > 0) || (diffP > 0 && diffQ < 0) {
				intersections++
			}
		}
	}
	return intersections
}
