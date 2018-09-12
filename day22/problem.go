package day22

import (
	"errors"
	"fmt"
	"strings"
)

// WordSet represents a set of possible words.
type WordSet map[string]struct{}

var errNoSentencePossible = errors.New("No sentence can be built")

// OriginalSentence returns the original sentence from a concatenated sentence.
func OriginalSentence(words WordSet, noSpaces string) (string, error) {
	if len(noSpaces) == 0 {
		return "", nil
	}
	for word := range words {
		if len(word) <= len(noSpaces) && word == noSpaces[:len(word)] {
			rest, err := OriginalSentence(words, noSpaces[len(word):])
			if err == nil {
				return strings.TrimSpace(fmt.Sprintf("%s %s", word, rest)), nil
			}
		}
	}
	return "", errNoSentencePossible
}
