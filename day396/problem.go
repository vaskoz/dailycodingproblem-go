package day396

// LongestPalindromicSubsequenceBrute returns the Longest
// Palindromic Subsequence using brute force.
func LongestPalindromicSubsequenceBrute(str string) int {
	runes := []rune(str)

	return lpsBrute(runes)
}

func lpsBrute(r []rune) int {
	switch l := len(r); {
	case l == 1:
		return 1
	case l == 2 && r[0] == r[1]:
		return 2
	case r[0] == r[l-1]:
		return lpsBrute(r[1:l-1]) + 2
	default:
		return max(
			lpsBrute(r[:l-1]),
			lpsBrute(r[1:l]),
		)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// LongestPalindromicSubsequenceDP returns the Longest
// Palindromic Subsequence using Dynamic Programming.
// Runs in O(N^2) space and time.
func LongestPalindromicSubsequenceDP(str string) int {
	runes := []rune(str)

	return lpsDP(runes)
}

func lpsDP(r []rune) int {
	n := len(r)
	subprob := make([][]int, n)

	for i := range subprob {
		subprob[i] = make([]int, n)
		subprob[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			switch j := i + length - 1; {
			case r[i] == r[j] && length == 2:
				subprob[i][j] = 2
			case r[i] == r[j]:
				subprob[i][j] = subprob[i+1][j-1] + 2
			default:
				subprob[i][j] = max(subprob[i][j-1], subprob[i+1][j])
			}
		}
	}

	return subprob[0][n-1]
}
