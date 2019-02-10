package day172

type trie struct {
	letters map[rune]*trie
}

func buildTrie(words []string) *trie {
	t := &trie{make(map[rune]*trie)}
	for _, word := range words {
		curr := t
		for _, r := range word {
			if _, exists := curr.letters[r]; !exists {
				curr.letters[r] = &trie{make(map[rune]*trie)}
			}
			curr = curr.letters[r]
		}
	}
	return t
}

func buildWordIndicies(s string, try *trie, length int) []wordIndex {
	var indicies []wordIndex
	for i := range s {
		curr := try
		for j := i; j < i+length && j < len(s); j++ {
			if _, found := curr.letters[rune(s[j])]; !found {
				break
			}
			curr = curr.letters[rune(s[j])]
		}
		if len(curr.letters) == 0 {
			indicies = append(indicies, wordIndex{s[i : i+length], i})
		}
	}
	return indicies
}

type wordIndex struct {
	word  string
	index int
}

// ConcatenatedSubstringIndicies returns the starting indicies of substrings
// composed of all the words provided.
// Runs in O(k*N) time where N is the length of 's' and 'k' is the length
// of each word in 'words'.
func ConcatenatedSubstringIndicies(s string, words []string) []int {
	try := buildTrie(words)
	pos := make(map[string]int, len(words))
	for i, word := range words {
		pos[word] = i
	}
	length := len(words[0])
	indicies := buildWordIndicies(s, try, length)
	var result []int
	for i, ind := range indicies {
		w := make(map[string]struct{}, len(words))
		w[ind.word] = struct{}{}
		next := ind.index + length
		for j := i + 1; j < len(indicies); j++ {
			if indicies[j].index < next {
				continue
			} else if _, found := w[indicies[j].word]; found || indicies[j].index > next {
				break
			} else {
				w[indicies[j].word] = struct{}{}
				next += length
			}
		}
		if len(w) == len(words) {
			result = append(result, ind.index)
		}
	}
	return result
}
