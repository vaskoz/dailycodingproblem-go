package day157

// IsPermutationPalindrome answers if the given string
// can be permutated into a palindrome.
// Runs in O(N) time and O(N) space.
func IsPermutationPalindrome(s string) bool {
	runeCount := make(map[rune]int)
	for _, r := range s {
		runeCount[r]++
	}
	var odds int
	for _, v := range runeCount {
		if v%2 == 1 {
			odds++
		}
	}
	return odds < 2
}
