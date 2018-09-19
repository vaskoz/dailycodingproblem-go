package day29

import (
	"fmt"
	"strconv"
	"strings"
)

// RunLengthEncoding takes a string of runes [A-Z] and encodes it using run-length.
func RunLengthEncoding(str string) string {
	var sb strings.Builder
	var count int
	var last rune
	for i, r := range str {
		if i == 0 {
			count++
			last = r
			continue
		} else if r == last {
			count++
		} else {
			sb.WriteString(fmt.Sprintf("%d%s", count, string(last)))
			count = 1
			last = r
		}
	}
	sb.WriteString(fmt.Sprintf("%d%s", count, string(last)))
	return sb.String()
}

// RunLengthDecoding takes a run-length encoding and returns the original string.
func RunLengthDecoding(str string) string {
	var sb strings.Builder
	var count int
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			sb.WriteString(strings.Repeat(string(r), count))
			count = 0
		} else {
			if count != 0 {
				count *= 10
			}
			digit, _ := strconv.Atoi(string(r))
			count += digit
		}
	}
	return sb.String()
}
