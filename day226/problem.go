package day226

// AlienLanguageOrder returns the correct order of letters
// of an alien language.
func AlienLanguageOrder(sortedWords []string) []rune {
	g := make(map[rune]map[rune]struct{})
	for i := 1; i < len(sortedWords); i++ {
		var firstDiff int
		for firstDiff = 0; sortedWords[i-1][firstDiff] == sortedWords[i][firstDiff]; firstDiff++ {
		}
		if _, exists := g[rune(sortedWords[i-1][firstDiff])]; !exists {
			g[rune(sortedWords[i-1][firstDiff])] = make(map[rune]struct{})
		}
		g[rune(sortedWords[i-1][firstDiff])][rune(sortedWords[i][firstDiff])] = struct{}{}
	}
	return kahnsTopologicalSort(g)
}

func kahnsTopologicalSort(g map[rune]map[rune]struct{}) []rune {
	var result []rune
	s := make(map[rune]struct{})
	for k := range g {
		s[k] = struct{}{}
	}
	for _, v := range g {
		for k := range v {
			delete(s, k)
		}
	}
	for len(s) != 0 {
		var n rune
		for k := range s {
			n = k
			break
		}
		delete(s, n)
		result = append(result, n)
		for m := range g[n] {
			delete(g[n], m)
			if !hasIncomingEdges(g, m) {
				s[m] = struct{}{}
			}
		}

	}
	return result
}

func hasIncomingEdges(g map[rune]map[rune]struct{}, m rune) bool {
	for _, to := range g {
		for k := range to {
			if k == m {
				return true
			}
		}
	}
	return false
}
