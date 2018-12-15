package day113

import "strings"

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
	return ""
}
