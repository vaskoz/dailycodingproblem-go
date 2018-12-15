package day113

import (
	"strings"
)

// ReverseWords reverses the word position in the string.
// Runs in O(N) time and requires O(N) additional space.
func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

// ReverseWordsInPlace reverses the word position in the string.
// Runs in O(N) time and requires O(1) additional space.
func ReverseWordsInPlace(str string) string {
	r := []rune(str)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	var start, end int
	for end <= len(r) {
		if end == len(r) {
			reverseInPlace(r, start, end)
		} else if r[end] == ' ' {
			reverseInPlace(r, start, end)
			start = end
			for r[start] == ' ' {
				start++
			}
			end = start - 1
		}
		end++
	}
	return string(r)
}

func reverseInPlace(r []rune, start, end int) {
	for i := start; i < (start+end)/2; i++ {
		r[i], r[end-1-(i-start)] = r[end-1-(i-start)], r[i]
	}
}
