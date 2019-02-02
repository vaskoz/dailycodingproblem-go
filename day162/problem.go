package day162

import "strings"

// Trie is a prefix search tree.
type Trie struct {
	freq map[rune]int
	ptr  map[rune]*Trie
}

// ShortestUniquePrefix returns the shortest unique prefix for each word
// amongst the words in the list.
func ShortestUniquePrefix(words []string) []string {
	root := &Trie{make(map[rune]int), make(map[rune]*Trie)}
	for _, word := range words {
		curr := root
		for _, r := range word {
			curr.freq[r]++
			if curr.ptr[r] == nil {
				curr.ptr[r] = &Trie{make(map[rune]int), make(map[rune]*Trie)}
			}
			curr = curr.ptr[r]
		}
	}
	results := make([]string, 0, len(words))
	for _, word := range words {
		var sb strings.Builder
		curr := root
		for _, r := range word {
			sb.WriteRune(r)
			if curr.freq[r] == 1 {
				results = append(results, sb.String())
				break
			}
			curr = curr.ptr[r]
		}
	}
	return results
}
