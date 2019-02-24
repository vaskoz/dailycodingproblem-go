package day181

// MinimumPartitionPalindrome splits a string into as
// few strings as possible such that each string is a palindrome.
func MinimumPartitionPalindrome(str string) []string {
	if len(str) == 0 {
		return nil
	}
	palindrome := make([][]bool, len(str))
	for i := range palindrome {
		palindrome[i] = make([]bool, len(str))
		palindrome[i][i] = true
	}
	for l := 2; l <= len(str); l++ {
		for i := 0; i < len(str)-l+1; i++ {
			j := i + l - 1
			if l == 2 {
				palindrome[i][j] = str[i] == str[j]
			} else {
				palindrome[i][j] = (str[i] == str[j]) && palindrome[i+1][j-1]
			}
		}
	}
	cuts := findCutsBFS(str, palindrome)
	result := make([]string, 0, len(cuts))
	prev := 0
	for _, cut := range cuts {
		result = append(result, str[prev:cut])
		prev = cut
	}
	return result
}

func findCutsBFS(s string, pal [][]bool) []int {
	var q [][]int
	for i, isPal := range pal[0] {
		if isPal {
			q = append(q, []int{i + 1})
		}
	}
	var cuts []int
	for len(q) != 0 {
		cuts, q = q[0], q[1:]
		start := cuts[len(cuts)-1]
		if len(s) == start {
			break
		}
		for i := start; i < len(pal[start]); i++ {
			if pal[start][i] {
				newCuts := append([]int{}, cuts...)
				newCuts = append(newCuts, i+1)
				q = append(q, newCuts)
			}
		}
	}
	return cuts
}
