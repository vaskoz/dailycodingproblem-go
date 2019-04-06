package day111

import "reflect"

// StartIndicesAnagrams returns all the starting indices (0-based)
// for the anagrams of W in the longer string S.
// Runs in O(N*M) time where N is len(word) and M is len(str).
func StartIndicesAnagrams(w string, s string) []int {
	freq := make(map[rune]int)
	for _, r := range w {
		freq[r]++
	}
	var result []int
	for i := range s {
		sub := make(map[rune]int)
		for j := i; j < len(s) && j < i+len(w); j++ {
			sub[rune(s[j])]++
		}
		if reflect.DeepEqual(sub, freq) {
			result = append(result, i)
		}
	}
	return result
}
