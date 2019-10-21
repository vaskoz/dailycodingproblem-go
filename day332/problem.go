package day332

// Pair is a set of the A and B values that satisfy
// the requirements below.
type Pair struct {
	A, B int
}

// PositiveIntegerPairs returns all candidates
// that satisfy the following conditions:
// a + b = M
// a XOR b = N
func PositiveIntegerPairs(m, n int) []Pair {
	var pairs []Pair

	for a := 0; a < (m/2)+1; a++ {
		b := m - a
		if a^b == n {
			pairs = append(pairs, Pair{A: a, B: b})
		}
	}

	return pairs
}
