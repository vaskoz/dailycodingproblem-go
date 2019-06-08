package day290

import "reflect"

// MinimizeQux returns the smallest Qux slice
// possible.
// Runs in O(N*K) where K is the number of elements that can removed.
// Uses O(2*N) extra space to avoid modifying the input.
func MinimizeQux(input []rune) []rune {
	prev := make([]rune, len(input))
	next := make([]rune, len(input))
	copy(next, input)
	for !reflect.DeepEqual(prev, next) {
		prev = make([]rune, len(next))
		copy(prev, next)
		for i := 1; i < len(next); i++ {
			if next[i] != next[i-1] {
				missing := 'R' + 'G' + 'B' - next[i] - next[i-1]
				tmp := []rune{missing}
				next = append(next[:i-1], append(tmp, next[i+1:]...)...)
				break
			}
		}
	}
	return next
}
