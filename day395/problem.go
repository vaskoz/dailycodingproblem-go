package day395

import "sort"

// GroupByAnagram returns a slice of slices of strings
// grouped by their anagrams.
func GroupByAnagram(words []string) [][]string {
	m := make(map[string][]string)
	pos := make(map[string]int)
	posI := 0

	for _, word := range words {
		ana := []rune(word)
		sort.Slice(ana, func(i, j int) bool {
			return ana[i] < ana[j]
		})

		anaKey := string(ana)

		if _, seen := pos[anaKey]; !seen {
			pos[anaKey] = posI
			posI++
		}

		m[anaKey] = append(m[anaKey], word)
	}

	res := make([][]string, len(m))

	for k, v := range m {
		res[pos[k]] = v
	}

	return res
}
