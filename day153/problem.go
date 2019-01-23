package day153

import (
	"strings"
)

// SmallestDistanceBetweenWords returns the smallest number of words
// between the two given words in the text provided.
// Returns -1 if one or both of the words don't exist in the text.
func SmallestDistanceBetweenWords(text, one, two string) int {
	words := strings.Split(text, " ")
	oneI := getIndex(words, one, 0)
	twoI := getIndex(words, two, 0)
	if oneI == -1 || twoI == -1 {
		return -1
	}
	smallest := len(words)
	for oneI != -1 && twoI != -1 {
		if dist := abs(oneI - twoI); dist < smallest {
			smallest = dist
		}
		if oneI < twoI {
			oneI = getIndex(words, one, oneI+1)
		} else {
			twoI = getIndex(words, two, twoI+1)
		}
	}
	return smallest - 1
}

func getIndex(s []string, target string, start int) int {
	for i := start; i < len(s); i++ {
		if s[i] == target {
			return i
		}
	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
