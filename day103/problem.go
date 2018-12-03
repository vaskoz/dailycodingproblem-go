package day103

// ShortestSubstringContains returns the shortest substring that contains
// all the letters provided in any order.
// Windowing runs in O(N) time and O(letters) extra space.
func ShortestSubstringContains(str string, letters []rune) string {
	occ := make(map[rune]int, len(letters))
	for _, r := range letters {
		occ[r]++
	}
	result := str
	var start int
	for end, r := range str {
		if _, found := occ[r]; found {
			occ[r]--
			if allFound(occ) {
				for {
					if count, found := occ[rune(str[start])]; found && count+1 > 0 {
						break
					} else if found {
						occ[rune(str[start])]++
					}
					start++
				}
				result = candidate(str[start:end+1], result)
			}
		}
	}
	return answer(result, !allFound(occ))
}

func answer(result string, notPossible bool) string {
	if notPossible {
		return ""
	}
	return result
}

func candidate(incumbent, challenger string) string {
	if len(incumbent) <= len(challenger) {
		return incumbent
	}
	return challenger
}

func allFound(m map[rune]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
