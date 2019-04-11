package day232

// PrefixMapSum represents a prefix tree collection of sums.
type PrefixMapSum interface {
	Insert(key string, value int)
	Sum(prefix string) int
}

type prefixMapSumTrie struct {
	sum  int
	next map[rune]*prefixMapSumTrie
}

// NewPrefixMapSum returns a new instance.
func NewPrefixMapSum() PrefixMapSum {
	return &prefixMapSumTrie{}
}

func (pms *prefixMapSumTrie) Insert(key string, value int) {
	curr := pms
	for _, r := range key {
		if curr.next == nil {
			curr.next = make(map[rune]*prefixMapSumTrie)
		}
		if _, found := curr.next[r]; !found {
			curr.next[r] = &prefixMapSumTrie{}
		}
		curr.next[r].sum += value
		curr = curr.next[r]
	}
}

func (pms *prefixMapSumTrie) Sum(prefix string) int {
	curr := pms
	for _, r := range prefix {
		if curr.next == nil || curr.next[r] == nil {
			curr = nil
			break
		}
		curr = curr.next[r]
	}
	if curr != nil {
		return curr.sum
	}
	return 0
}
