package day108

import (
	"strings"
)

// EqualWithShifting answers if A can be converted to B
// by shifting.
// Runs in O(N) time.
func EqualWithShifting(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	startLetter := rune(a[0])
	startPos := strings.IndexRune(b, startLetter)
	for {
		if startPos == -1 {
			return false
		}
		matches := 0
		for pos := startPos; pos < startPos+len(b); pos++ {
			if a[pos-startPos] != b[pos%len(b)] {
				newIndex := strings.IndexRune(b[startPos+1:], startLetter)
				if newIndex == -1 {
					startPos = -1
				} else {
					startPos += (newIndex + 1)
				}
				break
			}
			matches++
		}
		if matches == len(a) {
			return true
		}
	}
}
