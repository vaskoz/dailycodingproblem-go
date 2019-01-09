package day121

// IsKPalindromeBrute answers if the input can be made into a palindrome
// by at most k-deletions.
func IsKPalindromeBrute(str string, k int) bool {
	orig := []rune(str)
	rev := make([]rune, len(orig))
	for i := 0; i < len(orig); i++ {
		rev[i] = orig[len(orig)-1-i]
	}
	return isKPalindromeBrute(orig, rev, len(orig), len(rev)) <= k*2
}

func isKPalindromeBrute(orig, rev []rune, oi, ri int) int {
	if oi == 0 {
		return ri
	}
	if ri == 0 {
		return oi
	}
	if orig[oi-1] == rev[ri-1] {
		return isKPalindromeBrute(orig, rev, oi-1, ri-1)
	}
	return 1 +
		min(isKPalindromeBrute(orig, rev, oi-1, ri),
			isKPalindromeBrute(orig, rev, oi, ri-1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
