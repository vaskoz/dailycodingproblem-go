package day260

// Unjumble uses hints to unjumble the input and return the result.
func Unjumble(hints []rune) []int {
	result := make([]int, 0, len(hints)+1)
	count := 0
	for _, r := range hints {
		if r == '+' {
			count++
		}
	}
	first := len(hints) - count
	result = append(result, first)
	small, large := first-1, first+1
	for _, r := range hints {
		if r == '+' {
			result = append(result, large)
			large++
		} else {
			result = append(result, small)
			small--
		}
	}
	return result
}
