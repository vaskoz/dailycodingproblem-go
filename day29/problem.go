package day29

import (
	"strconv"
	"strings"
)

// RunLengthEncoding takes a string of runes [A-Z] and encodes it using run-length.
func RunLengthEncoding(str string) string {
	var sb strings.Builder
	var count int
	var last rune
	for i, r := range str {
		switch {
		case i == 0:
			count++
			last = r
			continue
		case r == last:
			count++
		default:
			sb.WriteString(strconv.Itoa(count)) // nolint: gosec
			sb.WriteString(string(last))        // nolint: gosec
			count = 1
			last = r
		}
	}
	sb.WriteString(strconv.Itoa(count)) // nolint: gosec
	sb.WriteString(string(last))        // nolint: gosec
	return sb.String()
}

// RunLengthDecoding takes a run-length encoding and returns the original string.
func RunLengthDecoding(str string) string {
	var sb strings.Builder
	var count int
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			sb.WriteString(strings.Repeat(string(r), count)) // nolint: gosec
			count = 0
		} else {
			if count != 0 {
				count *= 10
			}
			digit, _ := strconv.Atoi(string(r)) // nolint: gosec
			count += digit
		}
	}
	return sb.String()
}
