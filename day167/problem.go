package day167

// Pair represents indicies that correspond to words
// that when concatenated result in a palindrome.
type Pair struct {
	First, Second int
}

// PairsPalindromes returns a slice of Pairs.
// Runs in O(N^2) time.
func PairsPalindromes(words []string) []Pair {
	var result []Pair
	for i, first := range words {
		for j, second := range words {
			if i == j {
				continue
			}
			combined := first + second
			if isPalindrome(combined) {
				result = append(result, Pair{i, j})
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
