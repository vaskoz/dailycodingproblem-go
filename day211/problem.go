package day211

// StartingIndicies returns the starting indicies
// of all occurrences of the pattern in the str.
// Runs in O(N*K) time where N=len(str) and K=len(pattern).
func StartingIndicies(str, pattern string) []int {
	var result []int
	for i := range str {
		if len(pattern) > len(str)-i {
			break
		}
		complete := true
		for j, pr := range pattern {
			if pr != rune(str[i+j]) {
				complete = false
				break
			}
		}
		if complete {
			result = append(result, i)
		}
	}
	return result
}
