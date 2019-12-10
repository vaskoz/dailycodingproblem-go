package day386

import "sort"

// SortStringByCharFreq sorts the string by character
// frequency with the most frequent letters first.
// For a tie-breaker, they're sorted in reverse alphabetical
// order.
func SortStringByCharFreq(s string) string {
	freq := make(map[rune]int)
	runes := []rune(s)

	for _, r := range runes {
		freq[r]++
	}

	sort.Slice(runes, func(i, j int) bool {
		if f1, f2 := freq[runes[i]], freq[runes[j]]; f1 > f2 {
			return true
		} else if f1 == f2 {
			return runes[i] > runes[j]
		}

		return false
	})

	return string(runes)
}
