package day347

import "strings"

// LexicographicallySmallest returns the lexicographically smallest
// string that can be created after an unlimited number of moves
func LexicographicallySmallest(str string, k int) string {
	var sb strings.Builder
	result := str
	for len(str) > 0 {
		smallest := rune(str[0])
		smallestIndex := 0
		for i, r := range str {
			if i < k && r < smallest {
				smallest = r
				smallestIndex = i
			}
		}
		str = str[:smallestIndex] + str[smallestIndex+1:]
		sb.WriteString(string(smallest))
		if candidate := str + sb.String(); candidate < result {
			result = candidate
		}
	}
	return result
}
