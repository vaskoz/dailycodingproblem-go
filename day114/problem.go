package day114

import "strings"

// ReverseWordsMaintainDelimters reverses the words in the string
// while maintaining the relative order of the delimiters
// Works only with two delimiters "/" or ":"
// Runs in O(N) time and requires O(N) extra space to construct the result.
func ReverseWordsMaintainDelimters(str string) string {
	var start, end int
	var words []string
	var delimiters []string
	for end <= len(str) {
		if end == len(str) {
			if str[end-1] != ':' && str[end-1] != '/' {
				words = append(words, str[start:end])
			}
			break
		}
		if str[end] == '/' || str[end] == ':' {
			delimiters = append(delimiters, string(str[end]))
			words = append(words, str[start:end])
			start = end + 1
		}
		end++
	}
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	var result []string
	if len(delimiters) != len(words) {
		delimiters = append([]string{""}, delimiters...)
	}
	for i := range words {
		result = append(result, delimiters[i], words[i])
	}
	return strings.Join(result, "")
}
