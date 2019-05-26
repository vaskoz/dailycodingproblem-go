package day259

// GhostWinningLetterPlayerOne determine the letters the
// first player should start with, such that with optimal
// play they cannot lose.
func GhostWinningLetterPlayerOne(dict []string) []rune {
	firstRuneAllEven := make(map[rune]bool)
	for _, word := range dict {
		runes := []rune(word)
		isEvenLength := len(runes)%2 == 0
		firstRune := runes[0]
		if _, found := firstRuneAllEven[firstRune]; !isEvenLength {
			firstRuneAllEven[firstRune] = false
		} else if !found && isEvenLength {
			firstRuneAllEven[firstRune] = true
		}
	}
	var result []rune
	for r, valid := range firstRuneAllEven {
		if valid {
			result = append(result, r)
		}
	}
	return result
}
