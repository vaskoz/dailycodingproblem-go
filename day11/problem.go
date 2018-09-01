package day11

// Trie represents a prefix tree.
type Trie struct {
	letters   map[rune]*Trie
	wordCount int
}

// NewTrie constructs a new trie with the given words.
func NewTrie(words []string) *Trie {
	root := &Trie{letters: make(map[rune]*Trie)}
	for _, word := range words {
		current := root
		for _, letter := range word {
			if n, found := current.letters[letter]; !found {
				current.letters[letter] = &Trie{letters: make(map[rune]*Trie)}
				current = current.letters[letter]
			} else {
				current = n
			}
		}
		current.wordCount++
	}
	return root
}

// Match returns a list of words that match the given prefix.
func (t *Trie) Match(prefix string) []string {
	var result []string
	current := t
	for _, r := range prefix {
		n, found := current.letters[r]
		if !found {
			return result // prefix not found anywhere
		}
		current = n
	}
	return append(result, dfs(current, prefix)...)
}

func dfs(trie *Trie, prefix string) []string {
	var words []string
	if len(trie.letters) == 0 {
		return []string{prefix}
	}
	for letter, node := range trie.letters {
		words = append(words, dfs(node, prefix+string(letter))...)
	}
	if trie.wordCount > 0 {
		words = append(words, prefix)
	}
	return words
}
