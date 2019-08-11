package day345

import "strings"

// Thesaurus represents a thesaurus of words.
type Thesaurus map[string]map[string]struct{}

// AreSentencesEquivalentTransitive answers if the two sentences are
// equivalent given the Thesaurus.
func AreSentencesEquivalentTransitive(s1, s2 string, thes Thesaurus) bool {
	p1 := strings.Split(s1, " ")
	p2 := strings.Split(s2, " ")
	if len(p1) != len(p2) {
		return false
	}
	newThes := make(Thesaurus)
	for word := range thes {
		newThes[word] = make(map[string]struct{})
		words := make([]string, 0, len(thes[word]))
		for next := range thes[word] {
			words = append(words, next)
			if _, exists := newThes[next]; !exists {
				newThes[next] = make(map[string]struct{})
			}
			newThes[word][next] = struct{}{}
			newThes[next][word] = struct{}{}
		}
		for i := range words {
			for j := i + 1; j < len(words); j++ {
				newThes[words[i]][words[j]] = struct{}{}
				newThes[words[j]][words[i]] = struct{}{}
			}
		}
	}
	thes = newThes
	for i := range p1 {
		if p1[i] != p2[i] &&
			!isConvertible(p1[i], p2[i], thes, make(map[string]struct{})) {
			return false
		}
	}
	return true
}

// AreSentencesEquivalent answers if the two sentences are
// equivalent given the Thesaurus.
func AreSentencesEquivalent(s1, s2 string, thes Thesaurus) bool {
	p1 := strings.Split(s1, " ")
	p2 := strings.Split(s2, " ")
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if p1[i] != p2[i] &&
			!isConvertible(p1[i], p2[i], thes, make(map[string]struct{})) {
			return false
		}
	}
	return true
}

func isConvertible(src, target string, thes Thesaurus,
	visited map[string]struct{}) bool {
	if src == target {
		return true
	}
	visited[src] = struct{}{}
	for next := range thes[src] {
		if _, seen := visited[next]; !seen {
			if isConvertible(next, target, thes, visited) {
				return true
			}
		}
	}
	delete(visited, src)
	return false
}
