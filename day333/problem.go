package day333

// Knowing represents a personId knowing another personId.
// Return true of first person knows the second person.
// NOTE: This is not bidirectional.
type Knowing map[int]map[int]bool

// Knows is a O(1) method that answers if 'a' "knows" 'b'.
// This does not imply that 'b' "knows" 'a'.
func (k *Knowing) Knows(a, b int) bool {
	if known, exists := (*k)[a]; exists {
		return known[b]
	}

	return false
}

// FindCelebrity returns the ID of the celebrity.
// Runs in O(N) time and O(1) space.
func FindCelebrity(knows *Knowing, numPeople int) int {
	a := 0
	b := numPeople - 1

	for a < b {
		if knows.Knows(a, b) {
			a++
		} else {
			b--
		}
	}

	for i := 0; i < numPeople; i++ {
		if i != a && (knows.Knows(a, i) || !knows.Knows(i, a)) {
			return -1
		}
	}

	return a
}
