package day231

import (
	"sort"
	"strings"
)

// NoAdjacent returns a rearranged string
// where no adjacent characters are the same.
// Runs in O(N log N) due to sorting runes by frequency.
func NoAdjacent(s string) string {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	type frequency struct {
		r     rune
		count int
	}
	letters := make([]*frequency, 0, len(freq))
	for letter, count := range freq {
		letters = append(letters, &frequency{letter, count})
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i].count > letters[j].count
	})
	var sb strings.Builder
	for len(letters) != 0 {
		if len(letters) == 1 && letters[0].count > 1 {
			return ""
		}
		for _, f := range letters {
			sb.WriteRune(f.r)
			f.count--
		}
		for len(letters) != 0 && letters[len(letters)-1].count == 0 {
			letters = letters[:len(letters)-1]
		}
	}
	return sb.String()
}
