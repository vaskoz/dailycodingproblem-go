package day411

// IsKPalindromeBrute answers if the input can be made into a palindrome
// by at most k-deletions.
// Runtime is O(2^N).
func IsKPalindromeBrute(str string, k int) bool {
	orig := []rune(str)
	rev := make([]rune, len(orig))

	for i := 0; i < len(orig); i++ {
		rev[i] = orig[len(orig)-1-i]
	}

	return isKPalindromeBrute(orig, rev, len(orig), len(rev)) <= k*2
}

func isKPalindromeBrute(orig, rev []rune, oi, ri int) int {
	switch {
	case oi == 0:
		return ri
	case ri == 0:
		return oi
	case orig[oi-1] == rev[ri-1]:
		return isKPalindromeBrute(orig, rev, oi-1, ri-1)
	default:
		return 1 +
			min(isKPalindromeBrute(orig, rev, oi-1, ri),
				isKPalindromeBrute(orig, rev, oi, ri-1))
	}
}

// IsKPalindromeDP answers if the input can be made into a palindrome
// by at most k-deletions.
// Runtime and space are both O(N^2) where N is the length of the input.
func IsKPalindromeDP(str string, k int) bool {
	orig := []rune(str)
	rev := make([]rune, len(orig))

	for i := 0; i < len(orig); i++ {
		rev[i] = orig[len(orig)-1-i]
	}

	return isKPalindromeDP(orig, rev) <= k*2
}

func isKPalindromeDP(orig, rev []rune) int {
	dp := make([][]int, len(orig)+1)
	for i := range dp {
		dp[i] = make([]int, len(orig)+1)
		for j := range dp[i] {
			switch {
			case i == 0:
				dp[i][j] = j
			case j == 0:
				dp[i][j] = i
			case orig[i-1] == rev[j-1]:
				dp[i][j] = dp[i-1][j-1]
			default:
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(orig)][len(rev)]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
