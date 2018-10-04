package day28

import (
	"fmt"
	"strings"
)

// Justify returns a slice of justified strings.
func Justify(words []string, k int) []string {
	var line, result []string
	var lineLength int
	for i := 0; i < len(words); i++ {
		word := words[i]
		if lineLength+len(word)+len(line) <= k {
			lineLength += len(word)
			line = append(line, word)
		} else {
			result = append(result, buildLine(line, k, lineLength))
			line = line[:0] // truncate
			lineLength = 0
			i-- // retry this word
		}
	}
	last := buildLine(line, k, lineLength)
	if last != "" {
		result = append(result, last)
	}
	return result
}

func buildLine(line []string, k, lineLength int) string {
	if len(line) == 1 {
		spaces := k - lineLength
		return fmt.Sprintf("%s%s", line[0], strings.Repeat(" ", spaces))
	}
	var sb strings.Builder
	minimumSpaces := (k - lineLength) / len(line)
	extraSpaces := k - lineLength - ((len(line) - 1) * minimumSpaces)
	for i, word := range line {
		sb.WriteString(word) // nolint: gosec
		if i == len(line)-1 {
			break
		}
		spaces := minimumSpaces
		if i < extraSpaces {
			spaces++
		}
		sb.WriteString(strings.Repeat(" ", spaces)) // nolint: gosec
	}
	return sb.String()
}
