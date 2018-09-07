package day17

import (
	"strings"
)

// FindLongestAbsolutePathLength returns the length of the longest file path.
func FindLongestAbsolutePathLength(fs string) int {
	var parts []string
	var isFile bool
	var partStart, tabCount, maxLength int
	for i, r := range fs {
		if isNewStart(i, r, fs) {
			partStart = i
		} else if r == '.' {
			isFile = true
		} else if r == '\t' {
			tabCount++
		} else if r == '\n' {
			if pos := tabCount; len(parts) > pos {
				parts = parts[:pos]
			}
			parts = append(parts, fs[partStart:i])
			if len := len(strings.Join(parts, "/")); isFile && len > maxLength {
				maxLength = len
			}
			isFile = false
			tabCount = 0
		}
	}
	if isFile {
		parts = append(parts, fs[partStart:])
	}
	return len(strings.Join(parts, "/"))
}

func isNewStart(i int, r rune, fs string) bool {
	return i-1 > 0 && fs[i-1] == '\t' && r != '\t' && r != '\n'
}
