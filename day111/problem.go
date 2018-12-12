package day111

import "reflect"

// StartIndicesAnagrams returns all the starting indices (0-based)
// for the anagrams of W in the longer string S.
// Runs in O(N*M) time where N is len(word) and M is len(str).
func StartIndicesAnagrams(W string, S string) []int {
	freq := make(map[rune]int)
	for _, r := range W {
		freq[r]++
	}
	var result []int
	for i := range S {
		sub := make(map[rune]int)
		for j := i; j < len(S) && j < i+len(W); j++ {
			sub[rune(S[j])]++
		}
		if reflect.DeepEqual(sub, freq) {
			result = append(result, i)
		}
	}
	return result
}
