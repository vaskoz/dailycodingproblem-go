package day46

// IsPalindrome answers this question is linear time O(N) and O(1) space.
func IsPalindrome(letters []rune) bool {
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// LongestPalindromicSubstringBrute runs in O(N^2) time and O(1) space.
func LongestPalindromicSubstringBrute(str string) string {
	letters := []rune(str)
	for l := len(letters); l > 0; l-- {
		for left, right := 0, l; right <= len(letters); left, right = left+1, right+1 {
			if IsPalindrome(letters[left:right]) {
				return string(letters[left:right])
			}
		}
	}
	return ""
}
