package day246

// CircleWords returns a circle formation for these words.
// If a circle isn't possible, then returns nil.
func CircleWords(words []string) []string {
	if len(words) < 2 {
		return words
	}
	answer := make([]string, 0, len(words))
	answer = append(answer, words[0])
	remaining := make(map[rune]map[string]struct{}, len(words))
	for i := 1; i < len(words); i++ {
		word := words[i]
		firstRune := rune(word[0])
		if _, ok := remaining[firstRune]; !ok {
			remaining[firstRune] = make(map[string]struct{})
		}
		remaining[firstRune][word] = struct{}{}
	}
	return circleWords(answer, remaining)
}

func circleWords(answer []string, remaining map[rune]map[string]struct{}) []string {
	if len(remaining) == 0 {
		return answer
	}
	lastWord := answer[len(answer)-1]
	lastRune := rune(lastWord[len(lastWord)-1])
	for w := range remaining[lastRune] {
		delete(remaining[lastRune], w)
		if len(remaining[lastRune]) == 0 {
			delete(remaining, lastRune)
		}
		if result := circleWords(append(answer, w), remaining); result != nil {
			return result
		}
		if _, ok := remaining[lastRune]; !ok {
			remaining[lastRune] = make(map[string]struct{})
		}
		remaining[lastRune][w] = struct{}{}
	}
	return nil
}
