package day159

import "errors"

var errNoRecurringRune = errors.New("no recurring rune")

// ErrNoRecurringRune is the error returned when the input has no
// recurring runes. Every rune is unique.
func ErrNoRecurringRune() error {
	return errNoRecurringRune
}

// FirstRecurringRune returns the first rune that appears twice.
// If no rune is repeated, then an error is returned.
// Runs in O(N) time and potentially O(N) space.
func FirstRecurringRune(s string) (rune, error) {
	runes := make(map[rune]int)
	for _, r := range s {
		runes[r]++
		if count := runes[r]; count == 2 {
			return r, nil
		}
	}
	return 0, ErrNoRecurringRune()
}
