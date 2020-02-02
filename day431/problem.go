package day431

import "strings"

// ValidSentences parses text and returns the valid sentences only
// as defined by the problem.
func ValidSentences(text string) []string {
	sentences := make([]string, 0)
	start := 0

	spaceDelimited := strings.Split(text, " ")
	possibleSentence := false

	for i, part := range spaceDelimited {
		runes := []rune(part)
		if part == "" || !isValid(runes) {
			possibleSentence = false
			continue
		}

		if runes[0] >= 'A' && runes[0] <= 'Z' {
			possibleSentence = true
			start = i

			continue
		}

		if last := runes[len(runes)-1]; last == '.' || last == '!' || last == '?' {
			if possibleSentence {
				sentences = append(sentences, strings.Join(spaceDelimited[start:i+1], " "))
			}
		}
	}

	return sentences
}

func isValid(runes []rune) bool {
	first := runes[0]
	if !((first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z')) {
		return false
	} else if len(runes) == 1 {
		return true
	}

	for i := 1; i < len(runes)-1; i++ {
		if char := runes[i]; char < 'a' || char > 'z' {
			return false
		}
	}

	last := runes[len(runes)-1]

	return strings.ContainsRune(",.!?;:", last) || (last >= 'a' && last <= 'z')
}
