package day108

import (
	"strings"
)

// EqualWithShifting answers if A can be converted to B
// by shifting.
// Runs in O(N) time.
func EqualWithShifting(A, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if len(A) == 0 {
		return true
	}
	startLetter := rune(A[0])
	startPos := strings.IndexRune(B, startLetter)
	for {
		if startPos == -1 {
			return false
		}
		matches := 0
		for pos := startPos; pos < startPos+len(B); pos++ {
			if A[pos-startPos] != B[pos%len(B)] {
				newIndex := strings.IndexRune(B[startPos+1:], startLetter)
				if newIndex == -1 {
					startPos = -1
				} else {
					startPos += (newIndex + 1)
				}
				break
			}
			matches++
		}
		if matches == len(A) {
			return true
		}
	}
}
